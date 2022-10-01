package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type WalletsController = interfaces.IWalletsController
type walletsController struct {
	service services.WalletsService
}

func NewWalletsController(service services.WalletsService) *walletsController {
	return &walletsController{service: service}
}

/*
@depsription CreateWalletsController
*/

func (c *walletsController) CreateController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOWalletsCreate{}

	res := c.service.CreateService(r.Context(), &body)
	helpers.Send(rw, helpers.ApiResponse(res))
}

/*
@depsription GetWalletsByIdController
*/

func (c *walletsController) GetByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOWalletsById{Id: uint(Id)}

	res := c.service.GetByIdService(r.Context(), &params)
	helpers.Send(rw, helpers.ApiResponse(res))
}

/*
@depsription UpdateWalletsByIdController
*/

func (c *walletsController) UpdateByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOWalletsById{Id: uint(Id)}

	body := dtos.DTOWalletsUpdate{}
	json.NewDecoder(r.Body).Decode(&body)

	res := c.service.UpdateByIdService(r.Context(), &body, &params)
	helpers.Send(rw, helpers.ApiResponse(res))
}
