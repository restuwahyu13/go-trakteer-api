package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"github.com/restuwahyu13/go-trakteer-api/controllers"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type rolesRoute struct {
	controller controllers.RolesController
	prefix     string
	router     *chi.Mux
}

func NewRolesRoute(prefix string, db *sqlx.DB, router *chi.Mux) *rolesRoute {
	repository := repositorys.NewRolesRepository(db)
	service := services.NewRolesService(repository)
	controller := controllers.NewRolesController(service)

	return &rolesRoute{controller: controller, prefix: prefix, router: router}
}

func (ctx *rolesRoute) RolesRoute() {
	ctx.router.Post(helpers.Endpoint(ctx.prefix, "/"), ctx.controller.CreateController)
	ctx.router.Get(helpers.Endpoint(ctx.prefix, "/"), ctx.controller.GetAllController)
	ctx.router.Get(helpers.Endpoint(ctx.prefix, "/{id}"), ctx.controller.GetByIdController)
	ctx.router.Delete(helpers.Endpoint(ctx.prefix, "/{id}"), ctx.controller.DeleteByIdController)
	ctx.router.Put(helpers.Endpoint(ctx.prefix, "/{id}"), ctx.controller.UpdatedByIdController)
}
