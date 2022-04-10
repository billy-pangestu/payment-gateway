package bootstrap

import (
	"payment-gateway-backend/pkg/logruslogger"
	api "payment-gateway-backend/server/handler"
	"payment-gateway-backend/server/middleware"

	chimiddleware "github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"

	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

// RegisterRoutes ...
func (boot *Bootup) RegisterRoutes() {
	handlerType := api.Handler{
		DB:         boot.DB,
		EnvConfig:  boot.EnvConfig,
		Validate:   boot.Validator,
		Translator: boot.Translator,
		ContractUC: &boot.ContractUC,
		Jwe:        boot.Jwe,
		Jwt:        boot.Jwt,
	}
	mJwt := middleware.VerifyMiddlewareInit{
		ContractUC: &boot.ContractUC,
	}

	boot.R.Route("/v1", func(r chi.Router) {
		// Define a limit rate to 1000 requests per IP per request.
		rate, _ := limiter.NewRateFromFormatted("1000-S")
		store, _ := sredis.NewStoreWithOptions(boot.ContractUC.Redis, limiter.StoreOptions{
			Prefix:   "limiter_global",
			MaxRetry: 3,
		})
		rateMiddleware := stdlib.NewMiddleware(limiter.New(store, rate, limiter.WithTrustForwardHeader(true)))
		r.Use(rateMiddleware.Handler)

		// Logging setup
		r.Use(chimiddleware.RequestID)
		r.Use(logruslogger.NewStructuredLogger(boot.EnvConfig["LOG_FILE_PATH"], boot.EnvConfig["LOG_DEFAULT"]))
		r.Use(chimiddleware.Recoverer)

		// API
		r.Route("/api", func(r chi.Router) {
			userHandler := api.UserHandler{Handler: handlerType}
			r.Route("/auth", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Post("/login", userHandler.LoginHandler)
					r.Post("/register", userHandler.CreateHandler)
				})
				r.Group(func(r chi.Router) {
					r.Use(mJwt.VerifyJwtTokenCredential)
					r.Post("/logout", userHandler.LogoutHandler)
				})
			})

			transactionHandler := api.TransactionHandler{Handler: handlerType}
			r.Route("/payment", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(mJwt.VerifyJwtTokenCredential)
					r.Get("/id/", transactionHandler.GetByIDHandler)
					r.Get("/tag/", transactionHandler.GetByTagHandler)
					r.Get("/userid", transactionHandler.GetByUserIDHandler)
					r.Get("/total", transactionHandler.GetTotalAmountHandler)
					r.Post("/", transactionHandler.StoreHandler)
					r.Put("/id/", transactionHandler.UpdateHandler)
					r.Delete("/id/", transactionHandler.DeleteHandler)
				})
			})
		})
	})
}
