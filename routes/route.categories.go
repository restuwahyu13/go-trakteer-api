package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type categoriesRoute struct {
	controller controllers.CategoriesController
	prefix     string
	router     *chi.Mux
}

func NewCategoriesRoute(prefix string, db *sqlx.DB, router *chi.Mux) *categoriesRoute {
	repository := repositorys.NewCategoriesRepository(db)
	service := services.NewCategoriesService(repository)
	controller := controllers.NewCategoriesController(service)
	return &categoriesRoute{controller: controller, prefix: prefix, router: router}
}

func (ctx *categoriesRoute) CategoriesRoute() {
	ctx.router.Post(helpers.Endpoint(ctx.prefix, "/"), ctx.controller.CreateController)
	ctx.router.Get(helpers.Endpoint(ctx.prefix, "/"), ctx.controller.GetAllController)
	ctx.router.Get(helpers.Endpoint(ctx.prefix, "/{id}"), ctx.controller.GetByIdController)
	ctx.router.Delete(helpers.Endpoint(ctx.prefix, "/{id}"), ctx.controller.DeleteByIdController)
	ctx.router.Put(helpers.Endpoint(ctx.prefix, "/{id}"), ctx.controller.UpdatedByIdController)
}
