package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type usersRoute struct {
	controller controllers.UsersController
	prefix     string
	router     *chi.Mux
}

func NewUsersRoute(prefix string, db *sqlx.DB, router *chi.Mux) *usersRoute {
	repository := repositorys.NewUsersRepository(db)
	service := services.NewUsersService(repository)
	controller := controllers.NewUsersController(service)

	return &usersRoute{controller: controller, prefix: prefix, router: router}
}

func (ctx *usersRoute) UsersRoute() {
	ctx.router.Post(helpers.Endpoint(ctx.prefix, "/login"), ctx.controller.LoginController)
	ctx.router.Post(helpers.Endpoint(ctx.prefix, "/forgot-password"), ctx.controller.ForgotPasswordController)
	ctx.router.Post(helpers.Endpoint(ctx.prefix, "/reset-password"), ctx.controller.ResetPasswordController)
	ctx.router.Put(helpers.Endpoint(ctx.prefix, "/change-password"), ctx.controller.ChangePasswordController)
	ctx.router.Get(helpers.Endpoint(ctx.prefix, "/profile/{id}"), ctx.controller.GetProfileByIdController)
	ctx.router.Put(helpers.Endpoint(ctx.prefix, "/profile/{id}"), ctx.controller.UpdateProfileByIdController)
}
