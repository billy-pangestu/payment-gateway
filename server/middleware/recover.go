package middleware

// The original work was derived from Goji's middleware, source:
// https://github.com/zenazn/goji/tree/master/web/middleware

import (
	"fmt"
	"net/http"
	apiHandler "payment-gateway-backend/server/handler"
	"runtime/debug"

	chiMW "github.com/go-chi/chi/middleware"
)

// RecoverInit ...
type RecoverInit struct {
	Debug bool
}

// Recoverer is a middleware that recovers from panics, logs the panic (and a backtrace),
// and returns a HTTP 500 (Internal Server Error) status if
// possible. Recoverer prints a request ID if one is provided.
//
// Alternatively, look at https://github.com/pressly/lg middleware pkgs.
func (recInit RecoverInit) Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				logEntry := chiMW.GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rvr, debug.Stack())
				} else {
					debug.PrintStack()
				}

				msg := "Internal Server Error!"
				if recInit.Debug == true {
					msg = fmt.Sprintf("Panic: %v", rvr)
				}

				apiHandler.RespondWithJSON(w, 500, 500, msg, []map[string]interface{}{}, []map[string]interface{}{})
				return
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
