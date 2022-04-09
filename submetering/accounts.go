package submetering

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type PostalAddress struct {
	No         string
	Street     string
	City       string
	Province   string
	PostalCode string
}

type UserName struct {
	FirstName string
	LastName  string
}

type User struct {
	UUID          string
	Name          UserName
	EmailAddress  string
	AreaNumber    string
	PhoneNumber   string
	PostalAddress *PostalAddress
	PasswordHash  string
	Roles         []UserRole
	RegisteredAt  time.Time
	Status        UserStatus
}

type UserRole string

const (
	UserRoleAdmin  UserRole = "admin"
	UserRoleClient UserRole = "client"
)

type UserStatus string

const (
	UserStatusActive UserStatus = "active"
	UserStatusBanned UserStatus = "banned"
)

const (
	ScopeRefreshToken  = "refresh_token"
	ScopeAccessToken   = "access_token"
	ScopeResetPassword = "reset_password"
)

type AccountsService interface {
	// Login(ctx context.Context, req LoginRequest) (res *LoginResponse, err error)
	CreateUser(ctx context.Context, req CreateUserRequest) (res *User, err error)
	// ListUsers(ctx context.Context, filters ...ListUsersFilter) (res []*User, err error)
	// CountUsers(ctx context.Context, filters ...ListUsersFilter) (res uint64, err error)
	// GetUser(ctx context.Context, userUUID string) (res *User, err error)
	// UpdateUser(ctx context.Context, userUUID string, req UpdateUserRequest) (res *User, err error)
	// BanUser(ctx context.Context, userUUID string) (res *User, err error)
	// UnbanUser(ctx context.Context, userUUID string) (res *User, err error)
	// SendResetPasswordLink(ctx context.Context, req SendResetPasswordLinkRequest) (err error)
	// ResetPasswordByToken(ctx context.Context, req ResetPasswordByTokenRequest) (err error)
	// ChangePassword(ctx context.Context, userUUID string, req ChangePasswordRequest) (err error)
	GenerateToken(ctx context.Context, options ...GenerateTokenOption) (token *Token, err error)
	ValidateTokenString(ctx context.Context, tokenString string) (token *Token, err error)
	// RefreshToken(ctx context.Context, req RefreshTokenRequest) (res *LoginResponse, err error)
	ValidateBasicAuth(ctx context.Context, username, password string) (err error)
}

type RefreshTokenRequest struct {
	RefreshToken string
}

type LoginRequest struct {
	EmailAddress string
	Password     string
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
	User         *User
}

type CreateUserRequest struct {
	Name          UserName
	EmailAddress  string
	AreaNumber    string
	PhoneNumber   string
	PostalAddress *PostalAddress
	Roles         []UserRole
}

type UpdateUserRequest struct {
	Name          UserName
	EmailAddress  string
	AreaNumber    string
	PhoneNumber   string
	PostalAddress *PostalAddress
	Roles         []UserRole
}

type ResetPasswordByTokenRequest struct {
	Token                  string
	NewPassword            string
	IgnorePasswordStrength bool
}

type SendResetPasswordLinkRequest struct {
	EmailAddress string
	LinkURL      string
}

type ChangePasswordRequest struct {
	OldPassword            string
	NewPassword            string
	IgnorePasswordStrength bool
}

type ListUsersFilters struct {
	Offset uint64
	Limit  uint64
	Query  string
	Roles  []string
	Status UserStatus
}

type ListUsersFilter func(l *ListUsersFilters)

func ListUsersWithOffset(offset uint64) ListUsersFilter {
	return func(l *ListUsersFilters) {
		l.Offset = offset
	}
}

func ListUsersWithLimit(limit uint64) ListUsersFilter {
	return func(l *ListUsersFilters) {
		l.Limit = limit
	}
}

func ListUsersWithQuery(q string) ListUsersFilter {
	return func(l *ListUsersFilters) {
		l.Query = q
	}
}

func ListUsersWithRoles(roles ...string) ListUsersFilter {
	return func(l *ListUsersFilters) {
		l.Roles = roles
	}
}

func ListUsersWithStatus(status UserStatus) ListUsersFilter {
	return func(l *ListUsersFilters) {
		l.Status = status
	}
}

type GenerateTokenOption func(token *Token)

func SetExpirationTime(exp time.Time) GenerateTokenOption {
	return func(token *Token) {
		token.Token.SetExpirationTime(exp)
	}
}

func SetSubject(sub string) GenerateTokenOption {
	return func(token *Token) {
		token.Token.SetSubject(sub)
	}
}

func SetScopes(scopes ...string) GenerateTokenOption {
	return func(token *Token) {
		v := strings.Join(scopes, " ")
		token.Token.Set(claimScope, v)
	}
}

func Set(k string, v interface{}) GenerateTokenOption {
	return func(token *Token) {
		token.Token.Set(k, v)
	}
}

type UserByUUIDNotFoundError struct {
	UserUUID string
}

func (err UserByUUIDNotFoundError) Error() string {
	return fmt.Sprintf("user with uuid '%s' not found", err.UserUUID)
}

type UserByEmailAddressNotFoundError struct {
	EmailAddress string
}

func (err UserByEmailAddressNotFoundError) Error() string {
	return fmt.Sprintf("user with email address '%s' not found", err.EmailAddress)
}

type UserByEmailAddressAlreadyExistsError struct {
	EmailAddress string
}

func (err UserByEmailAddressAlreadyExistsError) Error() string {
	return fmt.Sprintf("user with email address '%s' already exists", err.EmailAddress)
}

type InvalidBasicAuthCredentialsError struct{}

func (err InvalidBasicAuthCredentialsError) Error() string {
	return "invalid username or password provided for basic auth"
}

type InvalidPasswordProvidedError struct{}

func (err InvalidPasswordProvidedError) Error() string {
	return "invalid password provided"
}

type UserIsBannedError struct {
	UserUUID string
}

func (err UserIsBannedError) Error() string {
	return fmt.Sprintf("user with uuid %s is banned", err.UserUUID)
}

type UserIsNotBannedError struct {
	UserUUID string
}

func (err UserIsNotBannedError) Error() string {
	return fmt.Sprintf("user with uuid %s is not banned", err.UserUUID)
}

type InvalidTokenProvidedError struct{}

func (err InvalidTokenProvidedError) Error() string {
	return "invalid token provided"
}

type NoPasswordSpecifiedError struct {
	UserUUID string
}

func (err NoPasswordSpecifiedError) Error() string {
	return fmt.Sprintf("no or bad password is specified for user with uuid '%s'", err.UserUUID)
}

type InvalidEmailAddressProvidedError struct {
	EmailAddress string
}

func (err InvalidEmailAddressProvidedError) Error() string {
	return fmt.Sprintf("email address '%s' is invalid", err.EmailAddress)
}

type InvalidLinkURLError struct {
	LinkURL string
}

func (err InvalidLinkURLError) Error() string {
	return fmt.Sprintf("link url '%s' is invalid", err.LinkURL)
}

type PoorPasswordError struct{}

func (err PoorPasswordError) Error() string {
	return "the password is not strong enough"
}
