package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/services"
)

type ControllerUsers struct {
	service *services.UsersService
}

func NewUsersController(service *services.UsersService) *ControllerUsers {
	return &ControllerUsers{service: service}
}

func (ctx *ControllerUsers) LoginController(rw http.ResponseWriter, r *http.Request) {
	req := dtos.DTOLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		rw.Write(helpers.ApiResponse(res))
	}

	res := ctx.service.LoginService(req)
	rw.Write(helpers.ApiResponse(res))
}
