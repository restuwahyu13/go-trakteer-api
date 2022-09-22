package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/middlewares"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type rolesRoute struct {
	controller controllers.RolesController
	prefix     string
	router     *chi.Mux
	db         *sqlx.DB
}

func NewRolesRoute(prefix string, db *sqlx.DB, router *chi.Mux) *rolesRoute {
	repository := repositorys.NewRolesRepository(db)
	service := services.NewRolesService(repository)
	controller := controllers.NewRolesController(service)

	return &rolesRoute{controller: controller, prefix: prefix, router: router, db: db}
}

func (r *rolesRoute) RolesRoute() {
	r.router.Route(r.prefix, func(router chi.Router) {
		router.Use(middlewares.NewMiddlewareAuth(r.db).Middleware)

		router.Post("/", r.controller.CreateController)
		router.Get("/", r.controller.GetAllController)
		router.Get("/{id:[0-9]+}", r.controller.GetByIdController)
		router.Delete("/{id:[0-9]+}", r.controller.DeleteByIdController)
		router.Put("/{id:[0-9]+}", r.controller.UpdatedByIdController)
	})
}
