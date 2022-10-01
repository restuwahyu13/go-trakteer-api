package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
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
	r.router.Route(r.prefix, func(route chi.Router) {
		route.Group(func(router chi.Router) {
			router.Post("/login", r.controller.LoginController)
			router.Post("/forgot-password", r.controller.ForgotPasswordController)
			router.Put("/reset-password/{token}", r.controller.ResetPasswordController)
			router.Put("/change-password/{id:[0-9]+}", r.controller.ChangePasswordController)
			router.Get("/profile/{id:[0-9]+}", r.controller.GetProfileByIdController)
			router.Put("/profile/{id:[0-9]+}", r.controller.UpdateProfileByIdController)
			router.Get("/health-token/{token}", r.controller.HealthCheckTokenController)
			router.Post("/refresh-token", r.controller.RefreshTokenController)
		})

		route.Group(func(router chi.Router) {
			router.Use(middlewares.NewMiddlewareAuth(r.db).Middleware)
			router.Use(middlewares.NewMiddlewarePermission("super admin", "staff").Middleware)

			router.Post("/", r.controller.CreateUsersController)
			router.Get("/", r.controller.GetAllUsersController)
			router.Get("/{id:[0-9]+}", r.controller.GetUsersByIdController)
			router.Delete("/{id:[0-9]+}", r.controller.DeleteUsersByIdController)
			router.Put("/{id:[0-9]+}", r.controller.UpdateUsersByIdController)
		})
	})
}
