package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type RolesController struct {
	service *services.RolesService
}

func NewRolesController(service *services.RolesService) *RolesController {
	return &RolesController{service: service}
}

/**
* @description CreateController
**/

func (ctx *RolesController) CreateController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.CreateService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}

/**
* @description GetAllController
**/

func (ctx *RolesController) GetAllController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.GetAllService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}

/**
* @description GetByIdController
**/

func (ctx *RolesController) GetByIdController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.GetByIdService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}

/**
* @description DeleteByIdController
**/

func (ctx *RolesController) DeleteByIdController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.DeleteByIdService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}

/**
* @description UpdatedByIdController
**/

func (ctx *RolesController) UpdatedByIdController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.UpdatedByIdService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}
