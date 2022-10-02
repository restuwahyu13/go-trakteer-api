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

type DonationTypeController = interfaces.IDonationTypeController
type donationTypeController struct {
	service services.DonationTypeService
}

func NewDonationTypeController(service services.DonationTypeService) *donationTypeController {
	return &donationTypeController{service: service}
}

/**
* @description CreateController
**/

func (c *donationTypeController) CreateController(rw http.ResponseWriter, r *http.Request) {
	body := dtos.DTODonationType{}
	iconRes, iconErr := helpers.FileUpload(r, "icon")

	if iconErr != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Uploading file Error: %s", iconErr)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	body.Name = r.FormValue("name")
	body.Icon = iconRes["icon"]

	// store original value
	name := body.Name

	if err := conform.Struct(r.Context(), &body); err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Go validator Error: %s", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	} else if errValidator := gpc.Validator(body); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	// assign normalize value
	body.Name = helpers.ReplaceAllString(`[^A-Za-z|\s]`, name, "")

	res := c.service.CreateService(r.Context(), &body)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description GetAllController
**/

func (c *donationTypeController) GetAllController(rw http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(helpers.QueryParser(r, "limit"))
	offset, _ := strconv.Atoi(helpers.QueryParser(r, "offset"))
	current_page, _ := strconv.Atoi(helpers.QueryParser(r, "current_page"))
	sort := strings.ToUpper(helpers.QueryParser(r, "sort"))

	queryOffset := uint(offset)
	query := dtos.DTODonationTypePagination{
		Limit:       uint(limit),
		Offset:      &queryOffset,
		Sort:        sort,
		CurrentPage: uint(current_page),
	}

	if errValidator := gpc.Validator(query); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := c.service.GetAllService(r.Context(), &query)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description GetByIdController
**/

func (c *donationTypeController) GetByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTODonationTypeId{Id: uint(Id)}

	if errValidator := gpc.Validator(params); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := c.service.GetByIdService(r.Context(), &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description DeleteByIdController
**/

func (c *donationTypeController) DeleteByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTODonationTypeId{Id: uint(Id)}

	if errValidator := gpc.Validator(params); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	res := c.service.DeleteByIdService(r.Context(), &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}

/**
* @description UpdatedByIdController
**/

func (c *donationTypeController) UpdatedByIdController(rw http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	params := dtos.DTODonationTypeId{Id: uint(Id)}

	body := dtos.DTODonationType{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Parse body to json error: %v", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	// store original value
	name := body.Name

	if err := conform.Struct(r.Context(), &body); err != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: fmt.Sprintf("Go validator Error: %s", err)}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	} else if errValidator := gpc.Validator(params); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	} else if errValidator := gpc.Validator(body); errValidator.Errors != nil {
		res := helpers.APIResponse{StatCode: http.StatusBadRequest, StatMsg: "Go validator Error", Data: errValidator.Errors}
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	// assign normalize value
	body.Name = helpers.ReplaceAllString(`[^A-Za-z|\s]`, name, "")

	res := c.service.UpdateByIdService(r.Context(), &body, &params)
	if res.StatCode >= 400 {
		helpers.Send(rw, helpers.ApiResponse(res))
		return
	}

	helpers.Send(rw, helpers.ApiResponse(res))
}
