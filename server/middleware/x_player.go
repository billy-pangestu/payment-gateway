package middleware

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

const (
	X_PLAYER = "X-PLAYER"
)

type Xplayer struct {
	DB *sqlx.DB
}

// RequestLoggerMiddleware ...
func (m Xplayer) RequestXplayerLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xp := r.Header.Get(X_PLAYER)
		if xp != "" {
			// xModel := model.NewXPlayerModel(m.DB)
			// _ = xModel.BrwStore()
		}

		next.ServeHTTP(w, r)
		return
	})
}
