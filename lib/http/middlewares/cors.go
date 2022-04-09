package middlewares

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func CORS() mux.MiddlewareFunc {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{
			http.MethodOptions,
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		}),
		handlers.AllowedHeaders([]string{
			"Authorization",
			"Content-Type",
			"Content-Language",
			"Accept",
			"Accept-Language",
			"Origin",
		}),
		handlers.AllowCredentials(),
		handlers.ExposedHeaders([]string{
			"Content-Type",
			"X-Pagination-Current-Page",
			"X-Pagination-Page-Count",
			"X-Pagination-Per-Page",
			"X-Pagination-Total-Count",
		}),
	)
}
