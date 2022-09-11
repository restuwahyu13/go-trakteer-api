package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

func CategoriesRoute(prefix string, db *sqlx.DB, router *chi.Mux) {
	repository := repositorys.NewCategoriesRepository(db)
	service := services.NewCategoriesService(repository)
	controller := controllers.NewCategoriesController(service)

	router.Post(helpers.Endpoint(prefix, "/"), controller.CreateController)
	router.Get(helpers.Endpoint(prefix, "/"), controller.GetAllController)
	router.Get(helpers.Endpoint(prefix, "/{id}"), controller.GetByIdController)
	router.Delete(helpers.Endpoint(prefix, "/{id}"), controller.DeleteByIdController)
	router.Put(helpers.Endpoint(prefix, "/{id}"), controller.UpdatedByIdController)
}
