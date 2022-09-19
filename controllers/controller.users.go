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

type UsersController = interfaces.IUsersController
type usersController struct {
	service services.UsersService
}

func NewUsersController(service services.UsersService) *usersController {
	return &usersController{service: service}
}

/**
* @description LoginController
**/

func (ctx *usersController) LoginController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOUsersLogin{}
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

	res := ctx.service.LoginService(&body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description ForgotPasswordController
**/

func (ctx *usersController) ForgotPasswordController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOUsersForgotPassword{}
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

	res := ctx.service.ForgotPasswordService(&body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description ResetPasswordController
**/

func (ctx *usersController) ResetPasswordController(rw http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	params := dtos.DTOUsersResetPasswordToken{Token: token}

	body := dtos.DTOUsersResetPassword{}
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

	res := ctx.service.ResetPasswordService(&body, &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description ChangePasswordController
**/

func (ctx *usersController) ChangePasswordController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOUsersById{Id: Id}

	body := dtos.DTOUsersChangePassword{}
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

	res := ctx.service.ChangePasswordService(&body, &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description GetProfileByIdController
**/

func (ctx *usersController) GetProfileByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOUsersGetProfileById{Id: Id}

	if errValidator := gpc.Validator(params); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Error Validators", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.GetProfileByIdService(&params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description UpdateProfileByIdController
**/

func (ctx *usersController) UpdateProfileByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOUsersGetProfileById{Id: Id}

	body := dtos.DTOUsersUpdateProfileById{}
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

	res := ctx.service.UpdateProfileByIdService(&body, &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description CreateUsersController
**/

func (ctx *usersController) CreateUsersController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOUsersCreate{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	if errValidator := gpc.Validator(body); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Error Validators", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.CreateUsersService(&body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description GetAllUsersController
**/

func (ctx *usersController) GetAllUsersController(rw http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(helpers.QueryParser(r, "limit"))
	offset, _ := strconv.Atoi(helpers.QueryParser(r, "offset"))
	current_page, _ := strconv.Atoi(helpers.QueryParser(r, "current_page"))
	sort := helpers.QueryParser(r, "sort")

	query := dtos.DTOUsersPagination{
		Limit:       limit,
		Offset:      offset,
		Sort:        strings.ToUpper(sort),
		CurrentPage: current_page,
	}

	if errValidator := gpc.Validator(query); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Error Validators", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.GetAllUsersService(&query)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description GetUsersByIdController
**/

func (ctx *usersController) GetUsersByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOUsersById{Id: Id}

	if errValidator := gpc.Validator(params); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Error Validators", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.GetUsersByIdService(&params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description DeleteUsersByIdController
**/

func (ctx *usersController) DeleteUsersByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOUsersById{Id: Id}

	if errValidator := gpc.Validator(params); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Error Validators", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.DeleteUsersByIdService(&params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description UpdateUsersByIdController
**/

func (ctx *usersController) UpdateUsersByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOUsersById{Id: Id}

	body := dtos.DTOUsersUpdate{}
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

	res := ctx.service.UpdateUsersByIdService(&body, &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}
