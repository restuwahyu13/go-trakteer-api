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
		router.Post(helpers.Endpoint(r.prefix, "/login"), r.controller.LoginController)
		router.Post(helpers.Endpoint(r.prefix, "/forgot-password"), r.controller.ForgotPasswordController)
		router.Put(helpers.Endpoint(r.prefix, "/reset-password/{token}"), r.controller.ResetPasswordController)
		router.Put(helpers.Endpoint(r.prefix, "/change-password/{id:[0-9]+}"), r.controller.ChangePasswordController)
		router.Get(helpers.Endpoint(r.prefix, "/profile/{id:[0-9]+}"), r.controller.GetProfileByIdController)
		router.Put(helpers.Endpoint(r.prefix, "/profile/{id:[0-9]+}"), r.controller.UpdateProfileByIdController)
	})

	r.router.Group(func(router chi.Router) {
		router.Use(middlewares.NewMiddlewareAuth(ctx.db).Middleware)
		router.Post(helpers.Endpoint(r.prefix, "/"), r.controller.CreateUsersController)
		router.Get(helpers.Endpoint(r.prefix, "/"), r.controller.GetAllUsersController)
		router.Get(helpers.Endpoint(r.prefix, "/{id:[0-9]+}"), r.controller.GetUsersByIdController)
		router.Delete(helpers.Endpoint(r.prefix, "/{id:[0-9]+}"), r.controller.DeleteUsersByIdController)
		router.Put(helpers.Endpoint(ctx.prefix, "/{id:[0-9]+}"), r.controller.UpdateUsersByIdController)
	})
}
