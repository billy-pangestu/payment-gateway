package middleware

import (
	"net/http"

	apiHandler "payment-gateway-backend/server/handler"

	"github.com/go-chi/chi"
)

// NotfoundMiddleware A custom not found response.
func NotfoundMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tctx := chi.NewRouteContext()
		rctx := chi.RouteContext(r.Context())

		if !rctx.Routes.Match(tctx, r.Method, r.URL.Path) {
			apiHandler.RespondWithJSON(w, 404, 404, "Request Not Found!", []map[string]interface{}{}, []map[string]interface{}{})
			return
		}

		next.ServeHTTP(w, r)
	})
}
