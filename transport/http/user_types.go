package http

import (
	"time"

	"github.com/himynamej/api-test.git/submetering"
	"github.com/nasermirzaei89/respond"
)

type RefreshTokenRequestBody struct {
	RefreshToken string `json:"refreshToken"`
}

type LoginRequestBody struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

type LoginSuccessResponseBody struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	RefreshToken string `json:"refreshToken,omitempty"`
	User         *User  `json:"user"`
}

func newLoginSuccessResponseBody(res *submetering.LoginResponse) *LoginSuccessResponseBody {
	if res == nil {
		return nil
	}

	rsp := new(LoginSuccessResponseBody)

	rsp.AccessToken = res.AccessToken
	rsp.TokenType = "bearer"
	rsp.RefreshToken = res.RefreshToken
	rsp.User = userFromEntity(res.User)

	return rsp
}

type PostalAddress struct {
	No         string `json:"no"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postalCode"`
}

type UserName struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type User struct {
	UUID          string         `json:"uuid"`
	Name          UserName       `json:"name"`
	EmailAddress  string         `json:"emailAddress"`
	AreaNumber    string         `json:"areaNumber"`
	PhoneNumber   string         `json:"phoneNumber"`
	PostalAddress *PostalAddress `json:"postalAddress,omitempty"`
	RegisteredAt  string         `json:"registeredAt"`
	Status        string         `json:"status"`
	Roles         []string       `json:"roles"`
}

func userFromEntity(user *submetering.User) *User {
	if user == nil {
		return nil
	}

	res := new(User)

	res.UUID = user.UUID
	res.Name.FirstName = user.Name.FirstName
	res.Name.LastName = user.Name.LastName
	res.EmailAddress = user.EmailAddress
	res.AreaNumber = user.AreaNumber
	res.PhoneNumber = user.PhoneNumber
	res.Status = string(user.Status)
	res.RegisteredAt = user.RegisteredAt.Format(time.RFC3339)

	if user.PostalAddress != nil {
		res.PostalAddress = &PostalAddress{
			No:         user.PostalAddress.No,
			Street:     user.PostalAddress.Street,
			City:       user.PostalAddress.City,
			Province:   user.PostalAddress.Province,
			PostalCode: user.PostalAddress.PostalCode,
		}
	}

	res.Roles = make([]string, len(user.Roles))
	for i := range user.Roles {
		res.Roles[i] = string(user.Roles[i])
	}

	return res
}

type CreateUserRequestBody struct {
	Name          UserName       `json:"name"`
	EmailAddress  string         `json:"emailAddress"`
	AreaNumber    string         `json:"areaNumber"`
	PhoneNumber   string         `json:"phoneNumber"`
	PostalAddress *PostalAddress `json:"postalAddress"`
	Roles         []string       `json:"roles"`
}

type CreateUserResponseBody struct {
	respond.WithStatusCreated
	UUID         string `json:"uuid"`
	RegisteredAt string `json:"registeredAt"`
}

// type UserList struct {
// 	Items      []*User    `json:"items"`
// 	Pagination Pagination `json:"pagination"`
// }

// func newUserList(users []*submetering.User, p Pagination) *UserList {
// 	res := new(UserList)

// 	res.Items = make([]*User, len(users))

// 	for i := range users {
// 		res.Items[i] = userFromEntity(users[i])
// 	}

// 	res.Pagination = p

// 	return res
// }

type UpdateUserRequestBody struct {
	Name          UserName       `json:"name"`
	EmailAddress  string         `json:"emailAddress"`
	AreaNumber    string         `json:"areaNumber"`
	PhoneNumber   string         `json:"phoneNumber"`
	PostalAddress *PostalAddress `json:"postalAddress"`
	Roles         []string       `json:"roles"`
}

type RequestPasswordResetRequestBody struct {
	EmailAddress string `json:"emailAddress"`
	LinkURL      string `json:"linkUrl"`
}

type ResetPasswordRequestBody struct {
	Token                  string `json:"token"`
	NewPassword            string `json:"newPassword"`
	IgnorePasswordStrength bool   `json:"ignorePasswordStrength"`
}

type ChangePasswordRequestBody struct {
	OldPassword            string `json:"oldPassword"`
	NewPassword            string `json:"newPassword"`
	IgnorePasswordStrength bool   `json:"ignorePasswordStrength"`
}
