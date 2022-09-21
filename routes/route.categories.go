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
	r.router.Group(func(router chi.Router) {
		router.Use(middlewares.NewMiddlewareAuth(r.db).Middleware)
		router.Post(helpers.Endpoint(r.prefix, "/"), r.controller.CreateController)
		router.Get(helpers.Endpoint(r.prefix, "/"), r.controller.GetAllController)
		router.Get(helpers.Endpoint(r.prefix, "/{id:[0-9]+}"), r.controller.GetByIdController)
		router.Delete(helpers.Endpoint(r.prefix, "/{id:[0-9]+}"), r.controller.DeleteByIdController)
		router.Put(helpers.Endpoint(r.prefix, "/{id:[0-9]+}"), r.controller.UpdatedByIdController)
	})
}
