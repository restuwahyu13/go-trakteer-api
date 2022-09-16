package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type CustomersController = interfaces.ICustomersController
type customersController struct {
	service services.CustomersService
}

func NewCustomersController(service services.CustomersService) *customersController {
	return &customersController{service: service}
}

/**
* @description RegisterController
**/

func (ctx *customersController) RegisterController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOCustomersRegister{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.RegisterService(&body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description LoginController
**/

func (ctx *customersController) LoginController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOCustomersLogin{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
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
* @description ActivationController
**/

func (ctx *customersController) ActivationController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOCustomersActivation{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.ActivationService(&body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description ResendActivationController
**/

func (ctx *customersController) ResendActivationController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOCustomersResendActivation{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.ResendActivationService(&body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description ForgotPasswordController
**/

func (ctx *customersController) ForgotPasswordController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOCustomersForgotPassword{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
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

func (ctx *customersController) ResetPasswordController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOCustomersResetPassword{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.ResetPasswordService(&body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description ChangePasswordController
**/

func (ctx *customersController) ChangePasswordController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOCustomersChangePassword{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := ctx.service.ChangePasswordService(&body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description GetProfileController
**/

func (ctx *customersController) GetProfileByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOCustomersGetProfileById{Id: Id}

	res := ctx.service.GetProfileByIdService(&params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description UpdateProfileController
**/

func (ctx *customersController) UpdateProfileByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOCustomersGetProfileById{Id: Id}

	body := dtos.DTOCustomersUpdateProfileById{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
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
