package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/middlewares"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type usersRoute struct {
	controller controllers.UsersController
	prefix     string
	router     *chi.Mux
	db         *sqlx.DB
}

func NewUsersRoute(prefix string, db *sqlx.DB, router *chi.Mux) *usersRoute {
	repository := repositorys.NewUsersRepository(db)
	service := services.NewUsersService(repository)
	controller := controllers.NewUsersController(service)

	return &usersRoute{controller: controller, prefix: prefix, router: router, db: db}
}

func (ctx *usersRoute) UsersRoute() {
	ctx.router.Group(func(r chi.Router) {
		r.Post(helpers.Endpoint(ctx.prefix, "/login"), ctx.controller.LoginController)
		r.Post(helpers.Endpoint(ctx.prefix, "/forgot-password"), ctx.controller.ForgotPasswordController)
		r.Put(helpers.Endpoint(ctx.prefix, "/reset-password/{token}"), ctx.controller.ResetPasswordController)
		r.Put(helpers.Endpoint(ctx.prefix, "/change-password/{id:[0-9]+}"), ctx.controller.ChangePasswordController)
		r.Get(helpers.Endpoint(ctx.prefix, "/profile/{id:[0-9]+}"), ctx.controller.GetProfileByIdController)
		r.Put(helpers.Endpoint(ctx.prefix, "/profile/{id:[0-9]+}"), ctx.controller.UpdateProfileByIdController)
	})

	ctx.router.Group(func(r chi.Router) {
		r.Use(middlewares.NewMiddlewareAuth(ctx.db).Middleware)
		r.Post(helpers.Endpoint(ctx.prefix, "/"), ctx.controller.CreateUsersController)
		r.Get(helpers.Endpoint(ctx.prefix, "/"), ctx.controller.GetAllUsersController)
		r.Get(helpers.Endpoint(ctx.prefix, "/{id:[0-9]+}"), ctx.controller.GetUsersByIdController)
		r.Delete(helpers.Endpoint(ctx.prefix, "/{id:[0-9]+}"), ctx.controller.DeleteUsersByIdController)
		r.Put(helpers.Endpoint(ctx.prefix, "/{id:[0-9]+}"), ctx.controller.UpdateUsersByIdController)
	})
}
