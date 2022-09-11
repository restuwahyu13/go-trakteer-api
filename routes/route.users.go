package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

func UsersRoute(prefix string, db *sqlx.DB, router *chi.Mux) {

	repository := repositorys.NewUsersRepository(db)
	service := services.NewUsersService(repository)
	controller := controllers.NewUsersController(service)

	router.Route(prefix, func(r chi.Router) {
		r.Post("/login", controller.LoginController)
	})
}
