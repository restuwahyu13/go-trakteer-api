package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type UsersController struct {
	service *services.UsersService
}

func NewUsersController(service *services.UsersService) *UsersController {
	return &UsersController{service: service}
}

/**
* @description RegisterController
**/

func (ctx *UsersController) RegisterController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.RegisterService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}

/**
* @description LoginController
**/

func (ctx *UsersController) LoginController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.LoginService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}

/**
* @description ActivationController
**/

func (ctx *UsersController) ActivationController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.ActivationService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}

/**
* @description ForgotPasswordController
**/

func (ctx *UsersController) ForgotPasswordController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.ForgotPasswordService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}

/**
* @description ResetPasswordController
**/

func (ctx *UsersController) ResetPasswordController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.ResetPasswordService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}

/**
* @description ChangePasswordController
**/

func (ctx *UsersController) ChangePasswordController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.ChangePasswordService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}

/**
* @description GetProfileController
**/

func (ctx *UsersController) GetProfileController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.GetProfileService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}

/**
* @description UpdateProfileController
**/

func (ctx *UsersController) UpdateProfileController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.UpdateProfileService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}
