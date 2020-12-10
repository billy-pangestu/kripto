package bootstrap

import (
	"xettle-backend/pkg/logruslogger"
	api "xettle-backend/server/handler"
	"xettle-backend/server/middleware"

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
			userAuthHandler := api.UserAuthHandler{Handler: handlerType}
			secretHandler := api.SecretHandler{Handler: handlerType}
			r.Route("/user", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(mJwt.VerifyJwtTokenCredential)
					r.Get("/token", userHandler.GetByTokenHandler)
					r.Post("/changepin", userHandler.ChangePin)
					r.Post("/verifypin", userHandler.VerifyPin)

				})
				r.Group(func(r chi.Router) {
					//test.php
					r.Post("/changepinwotoken", userHandler.SetPinWOToken)
					r.Post("/showepin", userHandler.ShowEpin)
					r.Post("/showppin", userHandler.ShowPpin)
					r.Post("/checkphone", userAuthHandler.CheckPhoneNumber)

					//home.php
					r.Get("/getnotes/{id}", secretHandler.GetNotes)
					r.Get("/getnote/{idnote}", secretHandler.GetNote)
					r.Get("/getdecriptnote/{idnote}/{password}", secretHandler.GetDecriptNote)
					r.Post("/insertnote", secretHandler.InsertNote)
					r.Put("/updatenote/{id}", secretHandler.UpdateNote)
					r.Delete("/deletenote/{id}", secretHandler.DeleteNote)

					//login.php || register.php
					r.Post("/login", userHandler.Login)
					r.Post("/register", userHandler.Register)
				})
			})
			otpHandler := api.OTPHandler{Handler: handlerType}
			r.Route("/otp", func(r chi.Router) {
				limit := middleware.LimitOTPInit{
					ContractUC:     &boot.ContractUC,
					MaxLimit:       50,
					Duration:       "24h",
					MaxLimitIDUser: 5,
				}

				r.Group(func(r chi.Router) {
					r.Use(limit.LimitByIP)
					r.Use(limit.LimitByID)
					r.Post("/request", otpHandler.RequestOTP)
				})
			})
			transactionHandler := api.TransactionHandler{Handler: handlerType}
			r.Route("/transaction", func(r chi.Router) {
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
			tagHandler := api.TagHandler{Handler: handlerType}
			r.Route("/tag", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(mJwt.VerifyJwtTokenCredential)
					r.Get("/id/", tagHandler.GetByIDHandler)
					r.Get("/userid", tagHandler.GetByUserIDHandler)
					r.Put("/id/", tagHandler.UpdateHandler)
					r.Delete("/id/", tagHandler.DeleteHandler)
				})
			})
			bankHandler := api.BankHandler{Handler: handlerType}
			r.Route("/bank", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(mJwt.VerifyJwtTokenCredential)
					r.Get("/id/", bankHandler.GetByIDHandler)
					r.Get("/userid", bankHandler.GetByUserIDHandler)
					r.Post("/", bankHandler.StoreHandler)
					r.Put("/id/", bankHandler.UpdateHandler)
					r.Delete("/id/", bankHandler.DeleteHandler)
				})
			})
		})

		// API ADMIN
		r.Route("/api-admin", func(r chi.Router) {
			adminHandler := api.AdminHandler{Handler: handlerType}
			r.Route("/admin", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Post("/login", adminHandler.LoginHandler)
				})
				r.Group(func(r chi.Router) {
					r.Use(mJwt.VerifySuperadminTokenCredential)
					r.Get("/", adminHandler.GetAllHandler)
					r.Get("/id/{id}", adminHandler.GetByIDHandler)
					r.Get("/code/{code}", adminHandler.GetByCodeHandler)
					r.Post("/", adminHandler.CreateHandler)
					r.Put("/id/{id}", adminHandler.UpdateHandler)
					r.Delete("/id/{id}", adminHandler.DeleteHandler)
				})
			})

			fileHandler := api.FileHandler{Handler: handlerType}
			r.Route("/file", func(r chi.Router) {
				r.Use(mJwt.VerifyAdminTokenCredential)
				r.Post("/", fileHandler.UploadHandler)
			})
		})
	})
}
