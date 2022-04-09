package accounts

import (
	"context"
	"fmt"
	"net/mail"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	repositories "github.com/himynamej/api-test.git/repository"
	"github.com/himynamej/api-test.git/submetering"

	"github.com/nasermirzaei89/jwt"
)

type Service struct {
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
	//mailingSvc                         submetering.MailingService
	passwordCost                       int
	issuer                             string
	signKey                            []byte
	verifyKey                          []byte
	algorithm                          jwt.Algorithm
	basicAuthUsername                  string
	basicAuthPassword                  string
	accessTokenLifetimeDuration        time.Duration
	refreshTokenLifetimeDuration       time.Duration
	resetPasswordTokenLifetimeDuration time.Duration
	clientPanelURL                     string
	minPasswordEntropyBits             float64
}

func (svc *Service) ValidateTokenString(_ context.Context, tokenString string) (*submetering.Token, error) {
	err := jwt.Verify(tokenString, svc.verifyKey)
	if err != nil {
		return nil, submetering.InvalidTokenProvidedError{}
	}

	t, err := jwt.Parse(tokenString)
	if err != nil {
		return nil, errors.Wrap(err, "error on parse token string")
	}

	err = t.Validate()
	if err != nil {
		return nil, submetering.InvalidTokenProvidedError{}
	}

	// TODO: check jti and rti in block list

	res := submetering.Token{
		SignKey: svc.signKey,
		Token:   t,
	}

	return &res, nil
}

func (svc *Service) ValidateBasicAuth(_ context.Context, username, password string) error {
	if username != svc.basicAuthUsername || password != svc.basicAuthPassword {
		return submetering.InvalidBasicAuthCredentialsError{}
	}

	return nil
}
func (svc *Service) GenerateToken(
	_ context.Context,
	options ...submetering.GenerateTokenOption,
) (*submetering.Token, error) {
	token := submetering.Token{
		SignKey: svc.signKey,
		Token:   jwt.New(svc.algorithm),
	}

	now := time.Now()

	token.Token.SetIssuedAt(now)
	token.Token.SetJWTID(uuid.NewString())

	if len(svc.issuer) > 0 {
		token.Token.SetIssuer(svc.issuer)
	}

	for i := range options {
		options[i](&token)
	}

	return &token, nil
}

func (svc *Service) CreateUser(ctx context.Context, req submetering.CreateUserRequest) (*submetering.User, error) {
	email, err := mail.ParseAddress(req.EmailAddress)
	if err != nil {
		return nil, submetering.InvalidEmailAddressProvidedError{EmailAddress: req.EmailAddress}
	}

	req.EmailAddress = email.Address

	_, err = svc.userRepo.FindByEmail(ctx, req.EmailAddress)
	if err != nil {
		if !errors.Is(err, repositories.ErrUserNotFound) {
			return nil, errors.Wrap(err, "error on find user by email address")
		}
	} else {
		return nil, submetering.UserByEmailAddressAlreadyExistsError{EmailAddress: req.EmailAddress}
	}

	user := submetering.User{
		UUID:          uuid.NewString(),
		Name:          req.Name,
		EmailAddress:  req.EmailAddress,
		AreaNumber:    req.AreaNumber,
		PhoneNumber:   req.PhoneNumber,
		PostalAddress: req.PostalAddress,
		PasswordHash:  "",
		Roles:         req.Roles,
		RegisteredAt:  time.Now(),
		Status:        submetering.UserStatusActive,
	}

	err = svc.userRepo.Insert(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "error on insert user to repository")
	}

	// send welcome email
	token, err := svc.GenerateToken(
		ctx,
		submetering.SetSubject(user.UUID),
		submetering.SetScopes(submetering.ScopeResetPassword),
	)
	if err != nil {
		return nil, errors.Wrap(err, "error on generate reset password token")
	}

	ts, err := token.Sign()
	if err != nil {
		return nil, errors.Wrap(err, "error on sign token")
	}

	tokenKey := strings.ReplaceAll(uuid.New().String(), "-", "")

	err = svc.tokenRepo.Set(ctx, tokenKey, ts, svc.resetPasswordTokenLifetimeDuration)
	if err != nil {
		return nil, errors.Wrap(err, "error on set token in repo")
	}

	u := fmt.Sprintf("%s/reset-password/%s", svc.clientPanelURL, tokenKey)

	_, err = url.ParseRequestURI(u)
	if err != nil {
		return nil, errors.Wrap(err, "error on parse request uri")
	}

	// body, err := mailtpl.RenderAccountCreated(u)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "error on render account created body")
	// }

	// err = svc.mailingSvc.SendMail(ctx, submetering.SendMailRequest{
	// 	To:      []string{user.EmailAddress},
	// 	Subject: "Account Created",
	// 	Body:    body,
	// })
	// if err != nil {
	// 	return nil, errors.Wrap(err, "error on send email")
	// }

	return &user, nil
}

const (
	DefaultAccessTokenLifetimeSeconds        int64   = 600      // 10 minutes
	DefaultRefreshTokenLifetimeSeconds       int64   = 31556952 // 1 year
	DefaultResetPasswordTokenLifetimeSeconds int64   = 86400    // 1 day
	MinPasswordEntropyBits                   float64 = 60
)

func New(
	userRepo repositories.UserRepository,
	tokenRepo repositories.TokenRepository,
	//mailingSvc submetering.MailingService,
	basicAuthUsername, basicAuthPassword string,
	clientPanelURL string,
	options ...Option,
) *Service {
	svc := Service{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		//mailingSvc:                         mailingSvc,
		issuer:                             "",
		signKey:                            []byte(""),
		verifyKey:                          []byte(""),
		algorithm:                          jwt.HS256,
		basicAuthUsername:                  basicAuthUsername,
		basicAuthPassword:                  basicAuthPassword,
		accessTokenLifetimeDuration:        time.Second * time.Duration(DefaultAccessTokenLifetimeSeconds),
		refreshTokenLifetimeDuration:       time.Second * time.Duration(DefaultRefreshTokenLifetimeSeconds),
		resetPasswordTokenLifetimeDuration: time.Second * time.Duration(DefaultResetPasswordTokenLifetimeSeconds),
		clientPanelURL:                     clientPanelURL,
		minPasswordEntropyBits:             MinPasswordEntropyBits,
	}

	for i := range options {
		options[i](&svc)
	}

	if svc.passwordCost > bcrypt.MaxCost || svc.passwordCost < bcrypt.MinCost {
		svc.passwordCost = bcrypt.DefaultCost
	}

	return &svc
}
