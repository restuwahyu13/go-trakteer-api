package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
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

func (ctx *customersRoute) CustomersRoute() {
	ctx.router.Post(helpers.Endpoint(ctx.prefix, "/register"), ctx.controller.RegisterController)
	ctx.router.Post(helpers.Endpoint(ctx.prefix, "/login"), ctx.controller.LoginController)
	ctx.router.Get(helpers.Endpoint(ctx.prefix, "/activation/{token}"), ctx.controller.ActivationController)
	ctx.router.Post(helpers.Endpoint(ctx.prefix, "/resend-activation"), ctx.controller.ResendActivationController)
	ctx.router.Post(helpers.Endpoint(ctx.prefix, "/forgot-password"), ctx.controller.ForgotPasswordController)
	ctx.router.Put(helpers.Endpoint(ctx.prefix, "/reset-password:{token}"), ctx.controller.ResetPasswordController)
	ctx.router.Put(helpers.Endpoint(ctx.prefix, "/change-password:[0-9]+"), ctx.controller.ChangePasswordController)
	ctx.router.Get(helpers.Endpoint(ctx.prefix, "/profile/{id:[0-9]+}"), ctx.controller.GetProfileByIdController)
	ctx.router.Put(helpers.Endpoint(ctx.prefix, "/profile/{id:[0-9]+}"), ctx.controller.UpdateProfileByIdController)
}
