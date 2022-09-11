package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

func UsersRoute(prefix string, db *sqlx.DB, router *chi.Mux) {

	repository := repositorys.NewUsersRepository(db)
	service := services.NewUsersService(repository)
	controller := controllers.NewUsersController(service)

	router.Post(helpers.Endpoint(prefix, "login"), controller.LoginController)
}
