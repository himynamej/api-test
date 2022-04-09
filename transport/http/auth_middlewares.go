package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/himynamej/api-test.git/lib/problem"
	"github.com/himynamej/api-test.git/submetering"
	"github.com/nasermirzaei89/respond"
)

type ContextKey string

const (
	ContextKeyUserUUID  ContextKey = "userUuid"
	ContextKeyUserRoles ContextKey = "userRoles"
)

func (h *handler) Authenticated() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenHeaderString := r.Header.Get("Authorization")
			if tokenHeaderString == "" {
				respond.Done(w, r, problem.Unauthorized("no authorization header provided"))

				return
			}

			var tokenString string

			const bearerPrefix = "Bearer "

			prefixLen := len(bearerPrefix)

			if !strings.HasPrefix(tokenHeaderString, bearerPrefix) {
				respond.Done(w, r, problem.Unauthorized("no authorization header provided"))
				fmt.Println("iiiiiiiiiiiiiiiiiiiiiin iffffffffffffffff")
				return
			}

			if len(tokenHeaderString) > prefixLen {
				tokenString = tokenHeaderString[prefixLen:]
			} else {
				respond.Done(w, r, problem.Unauthorized("invalid token"))

				return
			}

			token, err := h.accountsSvc.ValidateTokenString(r.Context(), tokenString)
			if err != nil {
				respond.Done(w, r, problem.Unauthorized(
					"invalid authorization header",
					problem.WithExtension("err", err.Error()),
				))

				return
			}

			if !token.HasScope(submetering.ScopeAccessToken) {
				respond.Done(w, r, problem.Unauthorized(
					"invalid authorization header",
					problem.WithExtension("err", accessTokenScopeNote),
				))

				return
			}

			r = r.WithContext(context.WithValue(r.Context(), ContextKeyUserUUID, token.MustGetSubject()))

			roles, err := token.Get("roles")
			if err == nil {
				r = r.WithContext(context.WithValue(r.Context(), ContextKeyUserRoles, interfaceToStringSlice(roles)))
			}

			next(w, r)
		}
	}
}

func interfaceToStringSlice(in interface{}) []string {
	vv, ok := in.([]interface{})
	if !ok {
		return []string{}
	}

	res := make([]string, 0)

	for i := range vv {
		v, ok := vv[i].(string)
		if !ok {
			continue
		}

		res = append(res, v)
	}

	return res
}

func getUserUUID(ctx context.Context) string {
	return ctx.Value(ContextKeyUserUUID).(string)
}

func (h *handler) HasRole(role submetering.UserRole) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			roles, ok := r.Context().Value(ContextKeyUserRoles).([]string)
			if ok {
				for i := range roles {
					if string(role) == roles[i] {
						next(w, r)

						return
					}
				}
			}

			respond.Done(w, r, problem.Forbidden("you are not allowed to do this action"))
		}
	}
}
