package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/middlewares"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type walletsRoute struct {
	controller controllers.WalletsController
	prefix     string
	router     *chi.Mux
	db         *sqlx.DB
}

func NewWalletsRoute(prefix string, db *sqlx.DB, router *chi.Mux) *walletsRoute {
	repository := repositorys.NewWalletsRepository(db)
	service := services.NewWalletsService(repository)
	controller := controllers.NewWalletsController(service)

	return &walletsRoute{controller: controller, prefix: prefix, router: router, db: db}
}

func (r *walletsRoute) WalletsRoute() {
	r.router.Route(r.prefix, func(route chi.Router) {
		route.Use(middlewares.NewMiddlewareAuth(r.db).Middleware)
		route.Use(middlewares.NewMiddlewarePermission("super admin", "staff", "customer").Middleware)

		route.Post("/", r.controller.CreateController)
		route.Get("/{id:[0-9]+}", r.controller.GetByIdController)
		route.Put("/{id:[0-9]+}", r.controller.UpdateByIdController)
	})
}
