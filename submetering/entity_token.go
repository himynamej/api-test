package submetering

import (
	"strings"

	"github.com/nasermirzaei89/jwt"
	"github.com/pkg/errors"
)

const (
	claimScope = "scope"
)

type Token struct {
	SignKey []byte
	Token   jwt.Token
}

func NewToken(signKey []byte, token jwt.Token) *Token {
	return &Token{
		SignKey: signKey,
		Token:   token,
	}
}

func (e Token) Sign() (string, error) {
	res, err := jwt.Sign(e.Token, e.SignKey)
	if err != nil {
		return "", errors.Wrap(err, "error on sign token")
	}

	return res, nil
}

func (e Token) GetID() (string, error) {
	res, err := e.Token.GetJWTID()
	if err != nil {
		return "", errors.Wrap(err, "error on get jwt id")
	}

	return res, nil
}

func (e Token) MustGetID() string {
	res, err := e.GetID()
	if err != nil {
		panic(err)
	}

	return res
}

func (e Token) GetSubject() (string, error) {
	res, err := e.Token.GetSubject()
	if err != nil {
		return "", errors.Wrap(err, "error on get jwt subject")
	}

	return res, nil
}

func (e Token) MustGetSubject() string {
	res, err := e.GetSubject()
	if err != nil {
		panic(err)
	}

	return res
}

func (e Token) Get(key string) (interface{}, error) {
	res, err := e.Token.Get(key)
	if err != nil {
		return nil, errors.Wrap(err, "error on get key")
	}

	return res, nil
}

func (e Token) HasScope(scope string) bool {
	res, err := e.Token.Get(claimScope)
	if err != nil {
		if errors.Is(err, jwt.ErrClaimNotFound) {
			return false
		}

		panic(errors.Wrap(err, "error on get token claim"))
	}

	scopesStr, ok := res.(string)
	if !ok {
		return false
	}

	scopes := strings.Split(scopesStr, " ")

	for i := range scopes {
		if scopes[i] == scope {
			return true
		}
	}

	return false
}
