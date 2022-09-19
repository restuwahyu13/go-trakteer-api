package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	gpc "github.com/restuwahyu13/go-playground-converter"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type RolesController = interfaces.IRolesController
type rolesController struct {
	service services.RolesService
}

func NewRolesController(service services.RolesService) *rolesController {
	return &rolesController{service: service}
}

/**
* @description CreateController
**/

func (ctx *rolesController) CreateController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTORoles{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	if errValidator := gpc.Validator(body); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Error Validators", Data: errValidator}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.CreateService(&body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description GetAllController
**/

func (ctx *rolesController) GetAllController(rw http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(helpers.QueryParser(r, "limit"))
	offset, _ := strconv.Atoi(helpers.QueryParser(r, "offset"))
	current_page, _ := strconv.Atoi(helpers.QueryParser(r, "current_page"))
	sort := helpers.QueryParser(r, "sort")

	query := dtos.DTORolePagination{
		Limit:       limit,
		Offset:      offset,
		Sort:        strings.ToUpper(sort),
		CurrentPage: current_page,
	}

	if errValidator := gpc.Validator(query); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Error Validators", Data: errValidator}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.GetAllService(&query)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description GetByIdController
**/

func (ctx *rolesController) GetByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTORolesById{Id: Id}

	if errValidator := gpc.Validator(params); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Error Validators", Data: errValidator}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.GetByIdService(&params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description DeleteByIdController
**/

func (ctx *rolesController) DeleteByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTORolesById{Id: Id}

	if errValidator := gpc.Validator(params); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Error Validators", Data: errValidator}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.DeleteByIdService(&params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description UpdatedByIdController
**/

func (ctx *rolesController) UpdatedByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTORolesById{Id: Id}

	body := dtos.DTORoles{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	if errValidator := gpc.Validator(params); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Error Validators", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	} else if errValidator := gpc.Validator(body); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Error Validators", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.UpdatedByIdService(&body, &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}
