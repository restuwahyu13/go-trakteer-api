package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	gpc "github.com/restuwahyu13/go-playground-converter"

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

func (c *customersController) RegisterController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOCustomersRegister{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	if errValidator := gpc.Validator(body); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := c.service.RegisterService(r.Context(), &body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description LoginController
**/

func (c *customersController) LoginController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOCustomersLogin{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	if errValidator := gpc.Validator(body); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := c.service.LoginService(r.Context(), &body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description ActivationController
**/

func (c *customersController) ActivationController(rw http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	params := dtos.DTOCustomersActivation{Token: token}

	if errValidator := gpc.Validator(params); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := c.service.ActivationService(r.Context(), &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description ResendActivationController
**/

func (c *customersController) ResendActivationController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOCustomersResendActivation{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	if errValidator := gpc.Validator(body); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := c.service.ResendActivationService(r.Context(), &body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description ForgotPasswordController
**/

func (c *customersController) ForgotPasswordController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTOCustomersForgotPassword{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	if errValidator := gpc.Validator(body); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := c.service.ForgotPasswordService(r.Context(), &body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description ResetPasswordController
**/

func (c *customersController) ResetPasswordController(rw http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	params := dtos.DTOCustomerResetPasswordToken{Token: token}

	body := dtos.DTOCustomersResetPassword{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	if errValidator := gpc.Validator(body); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := c.service.ResetPasswordService(r.Context(), &body, &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description ChangePasswordController
**/

func (c *customersController) ChangePasswordController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOCustomersById{Id: uint(Id)}

	body := dtos.DTOCustomersChangePassword{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	if errValidator := gpc.Validator(body); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := c.service.ChangePasswordService(r.Context(), &body, &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description GetProfileController
**/

func (c *customersController) GetProfileByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOCustomersGetProfileById{Id: uint(Id)}

	if errValidator := gpc.Validator(params); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := c.service.GetProfileByIdService(r.Context(), &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description UpdateProfileController
**/

func (c *customersController) UpdateProfileByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTOCustomersGetProfileById{Id: uint(Id)}

	body := dtos.DTOCustomersUpdateProfileById{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	if errValidator := gpc.Validator(params); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	} else if errValidator := gpc.Validator(body); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := c.service.UpdateProfileByIdService(r.Context(), &body, &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}
