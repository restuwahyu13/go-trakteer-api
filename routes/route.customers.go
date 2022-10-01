package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type customersRoute struct {
	controller controllers.CustomersController
	prefix     string
	router     *chi.Mux
}

func NewCustomersRoute(prefix string, db *sqlx.DB, router *chi.Mux) *customersRoute {
	repository := repositorys.NewCustomersRepository(db)
	service := services.NewCustomersService(repository)
	controller := controllers.NewCustomersController(service)

	return &customersRoute{controller: controller, prefix: prefix, router: router}
}

func (r *customersRoute) CustomersRoute() {
	r.router.Route(r.prefix, func(router chi.Router) {
		router.Post("/register", r.controller.RegisterController)
		router.Post("/login", r.controller.LoginController)
		router.Get("/activation/{token}", r.controller.ActivationController)
		router.Post("/resend-activation", r.controller.ResendActivationController)
		router.Post("/forgot-password", r.controller.ForgotPasswordController)
		router.Put("/reset-password:{token}", r.controller.ResetPasswordController)
		router.Put("/change-password:[0-9]+", r.controller.ChangePasswordController)
		router.Get("/profile/{id:[0-9]+}", r.controller.GetProfileByIdController)
		router.Put("/profile/{id:[0-9]+}", r.controller.UpdateProfileByIdController)
		router.Get("/health-token/{token}", r.controller.HealthCheckTokenController)
		router.Post("/refresh-token", r.controller.RefreshTokenController)
	})
}
