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
		r.router.Post("/register", r.controller.RegisterController)
		r.router.Post("/login", r.controller.LoginController)
		r.router.Get("/activation/{token}", r.controller.ActivationController)
		r.router.Post("/resend-activation", r.controller.ResendActivationController)
		r.router.Post("/forgot-password", r.controller.ForgotPasswordController)
		r.router.Put("/reset-password:{token}", r.controller.ResetPasswordController)
		r.router.Put("/change-password:[0-9]+", r.controller.ChangePasswordController)
		r.router.Get("/profile/{id:[0-9]+}", r.controller.GetProfileByIdController)
		r.router.Put("/profile/{id:[0-9]+}", r.controller.UpdateProfileByIdController)
	})
}
