package accounts

import (
	"time"

	"github.com/nasermirzaei89/jwt"
)

type Option func(svc *Service)

func SetPasswordCost(cost int) Option {
	return func(svc *Service) {
		svc.passwordCost = cost
	}
}

func WithIssuer(issuer string) Option {
	return func(svc *Service) {
		svc.issuer = issuer
	}
}

func WithRS256(privateKey, publicKey []byte) Option {
	return func(svc *Service) {
		svc.algorithm = jwt.RS256
		svc.signKey = privateKey
		svc.verifyKey = publicKey
	}
}

func WithAccessTokenLifetime(d time.Duration) Option {
	return func(svc *Service) {
		svc.accessTokenLifetimeDuration = d
	}
}

func SetMinPasswordEntropyBits(minPasswordEntropyBits float64) Option {
	return func(svc *Service) {
		svc.minPasswordEntropyBits = minPasswordEntropyBits
	}
}
