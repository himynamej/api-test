package middlewares

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/himynamej/api-test.git/lib/problem"
	"github.com/nasermirzaei89/respond"
	"github.com/pkg/errors"
)

type recoverMW struct {
	next http.Handler
}

func (mw *recoverMW) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if v := recover(); v != nil {
			err := errors.Errorf("panic recovered: %+v", v)
			rsp := problem.InternalServerError(err)
			respond.Done(w, r, rsp)
		}
	}()
	mw.next.ServeHTTP(w, r)
}

func RecoverPanic() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return &recoverMW{
			next: next,
		}
	}
}
