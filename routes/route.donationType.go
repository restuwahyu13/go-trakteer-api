package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/middlewares"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type donationTypeRoute struct {
	controller controllers.DonationTypeController
	prefix     string
	router     *chi.Mux
	db         *sqlx.DB
}

func NewDonationTypeRoute(prefix string, db *sqlx.DB, router *chi.Mux) *donationTypeRoute {
	repository := repositorys.NewDonationTypeRepository(db)
	service := services.NewDonationTypeService(repository)
	controller := controllers.NewDonationTypeController(service)
	return &donationTypeRoute{controller: controller, prefix: prefix, router: router, db: db}
}

func (r *donationTypeRoute) DonationTypeRoute() {
	r.router.Route(r.prefix, func(route chi.Router) {
		route.Use(middlewares.NewMiddlewareAuth(r.db).Middleware)
		route.Use(middlewares.NewMiddlewarePermission("super admin", "staff", "customer").Middleware)

		route.Post("/", r.controller.CreateController)
		route.Get("/", r.controller.GetAllController)
		route.Get("/{id:[0-9]+}", r.controller.GetByIdController)
		route.Delete("/{id:[0-9]+}", r.controller.DeleteByIdController)
		route.Put("/{id:[0-9]+}", r.controller.UpdatedByIdController)
	})
}
