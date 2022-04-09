package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/himynamej/api-test.git/lib/problem"
	"github.com/himynamej/api-test.git/submetering"
	"github.com/nasermirzaei89/respond"
	"github.com/pkg/errors"
)

func (h *handler) HandleCreateUser() http.HandlerFunc {
	hf := func(w http.ResponseWriter, r *http.Request) {
		var req CreateUserRequestBody

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			respond.Done(w, r, problem.BadRequest("invalid request body"))

			return
		}

		roles := make([]submetering.UserRole, len(req.Roles))

		name := submetering.UserName{FirstName: req.Name.FirstName, LastName: req.Name.LastName}

		for i := range req.Roles {
			roles[i] = submetering.UserRole(req.Roles[i])
		}

		res, err := h.accountsSvc.CreateUser(r.Context(), submetering.CreateUserRequest{
			Name:         name,
			EmailAddress: req.EmailAddress,
			AreaNumber:   req.AreaNumber,
			PhoneNumber:  req.PhoneNumber,
			PostalAddress: func() *submetering.PostalAddress {
				if req.PostalAddress != nil {
					return &submetering.PostalAddress{
						No:         req.PostalAddress.No,
						Street:     req.PostalAddress.Street,
						City:       req.PostalAddress.City,
						Province:   req.PostalAddress.Province,
						PostalCode: req.PostalAddress.PostalCode,
					}
				}

				return nil
			}(),
			Roles: roles,
		})
		if err != nil {
			var (
				err1 submetering.UserByEmailAddressAlreadyExistsError
				err2 submetering.UserByEmailAddressAlreadyExistsError
			)

			switch {
			case errors.As(err, &err1):
				respond.Done(w, r, problem.Conflict(err.Error()))
			case errors.As(err, &err2):
				respond.Done(w, r, problem.BadRequest(err.Error(), problem.WithExtension(noteKey, emailAddressNote)))
			default:
				respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on create user")))
			}

			return
		}

		rsp := CreateUserResponseBody{
			UUID:         res.UUID,
			RegisteredAt: res.RegisteredAt.Format(time.RFC3339),
		}

		respond.Done(w, r, rsp)
	}

	hf = h.HasRole(submetering.UserRoleAdmin)(hf)
	hf = h.Authenticated()(hf)

	return hf
}
