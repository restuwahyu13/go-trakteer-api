package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

func RolesRoute(prefix string, db *sqlx.DB, router *chi.Mux) {
	repository := repositorys.NewRolesRepository(db)
	service := services.NewRolesService(repository)
	controller := controllers.NewRolesController(service)

	router.Post(helpers.Endpoint(prefix, "/"), controller.CreateController)
	router.Get(helpers.Endpoint(prefix, "/"), controller.GetAllController)
	router.Get(helpers.Endpoint(prefix, "/:id"), controller.GetByIdController)
	router.Delete(helpers.Endpoint(prefix, "/:id"), controller.DeleteByIdController)
	router.Put(helpers.Endpoint(prefix, "/:id"), controller.UpdatedByIdController)
}
