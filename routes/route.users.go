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

	router.Post(helpers.Endpoint(prefix, "/register"), controller.RegisterController)
	router.Post(helpers.Endpoint(prefix, "/login"), controller.LoginController)
	router.Get(helpers.Endpoint(prefix, "/activation/:token"), controller.ActivationController)
	router.Post(helpers.Endpoint(prefix, "/forgot-password"), controller.ForgotPasswordController)
	router.Post(helpers.Endpoint(prefix, "/reset-password"), controller.ResetPasswordController)
	router.Put(helpers.Endpoint(prefix, "/change-password"), controller.ChangePasswordController)
	router.Get(helpers.Endpoint(prefix, "/profile/{id}"), controller.GetProfileController)
	router.Put(helpers.Endpoint(prefix, "/profile/{id}"), controller.UpdateProfileController)
}
