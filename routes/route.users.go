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

func (r *usersRoute) UsersRoute() {
	ctx.router.Group(func(router chi.Router) {
		router.Post(helpers.Endpoint(ctx.prefix, "/login"), ctx.controller.LoginController)
		router.Post(helpers.Endpoint(ctx.prefix, "/forgot-password"), ctx.controller.ForgotPasswordController)
		router.Put(helpers.Endpoint(ctx.prefix, "/reset-password/{token}"), ctx.controller.ResetPasswordController)
		router.Put(helpers.Endpoint(ctx.prefix, "/change-password/{id:[0-9]+}"), ctx.controller.ChangePasswordController)
		router.Get(helpers.Endpoint(ctx.prefix, "/profile/{id:[0-9]+}"), ctx.controller.GetProfileByIdController)
		router.Put(helpers.Endpoint(ctx.prefix, "/profile/{id:[0-9]+}"), ctx.controller.UpdateProfileByIdController)
	})

	r.router.Group(func(router chi.Router) {
		router.Use(middlewares.NewMiddlewareAuth(ctx.db).Middleware)
		router.Post(helpers.Endpoint(ctx.prefix, "/"), ctx.controller.CreateUsersController)
		router.Get(helpers.Endpoint(ctx.prefix, "/"), ctx.controller.GetAllUsersController)
		router.Get(helpers.Endpoint(ctx.prefix, "/{id:[0-9]+}"), ctx.controller.GetUsersByIdController)
		router.Delete(helpers.Endpoint(ctx.prefix, "/{id:[0-9]+}"), ctx.controller.DeleteUsersByIdController)
		router.Put(helpers.Endpoint(ctx.prefix, "/{id:[0-9]+}"), ctx.controller.UpdateUsersByIdController)
	})
}
