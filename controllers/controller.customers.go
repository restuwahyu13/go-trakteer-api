package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type CustomersController struct {
	service *services.CustomersService
}

func NewCustomersController(service *services.CustomersService) *CustomersController {
	return &CustomersController{service: service}
}

/**
* @description RegisterController
**/

func (ctx *CustomersController) RegisterController(rw http.ResponseWriter, r *http.Request) {
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

func (ctx *CustomersController) LoginController(rw http.ResponseWriter, r *http.Request) {
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

func (ctx *CustomersController) ActivationController(rw http.ResponseWriter, r *http.Request) {
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
* @description ResendActivationController
**/

func (ctx *CustomersController) ResendActivationController(rw http.ResponseWriter, r *http.Request) {
	apiResponse := make(chan helpers.APIResponse, 1)
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.ResendActivationService(req)
	apiResponse <- res

	helpers.Send(rw, helpers.ApiResponse(<-apiResponse))
}

/**
* @description ForgotPasswordController
**/

func (ctx *CustomersController) ForgotPasswordController(rw http.ResponseWriter, r *http.Request) {
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

func (ctx *CustomersController) ResetPasswordController(rw http.ResponseWriter, r *http.Request) {
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

func (ctx *CustomersController) ChangePasswordController(rw http.ResponseWriter, r *http.Request) {
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

func (ctx *CustomersController) GetProfileController(rw http.ResponseWriter, r *http.Request) {
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

func (ctx *CustomersController) UpdateProfileController(rw http.ResponseWriter, r *http.Request) {
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
