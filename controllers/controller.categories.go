package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type CategoriesController struct {
	service *services.CategoriesService
}

func NewCategoriesController(service *services.CategoriesService) *CategoriesController {
	return &CategoriesController{service: service}
}

/**
* @description CreateController
**/

func (ctx *CategoriesController) CreateController(rw http.ResponseWriter, r *http.Request) {
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.CreateService(req)
	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description GetAllController
**/

func (ctx *CategoriesController) GetAllController(rw http.ResponseWriter, r *http.Request) {
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.GetAllService(req)
	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description GetByIdController
**/

func (ctx *CategoriesController) GetByIdController(rw http.ResponseWriter, r *http.Request) {
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.GetByIdService(req)
	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description DeleteByIdController
**/

func (ctx *CategoriesController) DeleteByIdController(rw http.ResponseWriter, r *http.Request) {
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.DeleteByIdService(req)
	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description UpdatedByIdController
**/

func (ctx *CategoriesController) UpdatedByIdController(rw http.ResponseWriter, r *http.Request) {
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.UpdatedByIdService(req)
	helpers.Send(rw, helpers.ApiResponse(res))
}
