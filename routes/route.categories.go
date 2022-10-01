package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/middlewares"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type categoriesRoute struct {
	controller controllers.CategoriesController
	prefix     string
	router     *chi.Mux
	db         *sqlx.DB
}

func NewCategoriesRoute(prefix string, db *sqlx.DB, router *chi.Mux) *categoriesRoute {
	repository := repositorys.NewCategoriesRepository(db)
	service := services.NewCategoriesService(repository)
	controller := controllers.NewCategoriesController(service)
	return &categoriesRoute{controller: controller, prefix: prefix, router: router, db: db}
}

func (r *categoriesRoute) CategoriesRoute() {
	r.router.Route(r.prefix, func(route chi.Router) {
		route.Use(middlewares.NewMiddlewareAuth(r.db).Middleware)
		route.Use(middlewares.NewMiddlewarePermission("super admin", "staff").Middleware)

		route.Post("/", r.controller.CreateController)
		route.Get("/", r.controller.GetAllController)
		route.Get("/{id:[0-9]+}", r.controller.GetByIdController)
		route.Delete("/{id:[0-9]+}", r.controller.DeleteByIdController)
		route.Put("/{id:[0-9]+}", r.controller.UpdatedByIdController)
	})
}
