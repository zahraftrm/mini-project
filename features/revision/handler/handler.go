package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/zahraftrm/mini-project/app/middlewares"
	"github.com/zahraftrm/mini-project/features/revision"
	"github.com/zahraftrm/mini-project/helper"

	"github.com/labstack/echo/v4"
)

type RevisionHandler struct {
	revisionService revision.RevisionServiceInterface
}

func New(service revision.RevisionServiceInterface) *RevisionHandler {
	return &RevisionHandler{
		revisionService: service,
	}
}

func (handler *RevisionHandler) Create(c echo.Context) error {
	idToken,role := middlewares.ExtractTokenInfo(c, "admin")
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	dataInput := RevisionRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&dataInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	// mapping dari request ke core
	dataCore := revision.Core{
		TrainingID: dataInput.TrainingID,
		Message:  dataInput.Message,
	}
	err := handler.revisionService.Create(dataCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
}

func (handler *RevisionHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var idConv int
	if id != "" {
		var errConv error
		idConv, errConv = strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("param id is required and must be number"))
		}
	}

	// TODO : get ID from token
	idToken,role := middlewares.ExtractTokenInfo(c, "admin")
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	dataReq := RevisionRequest{}
	errBind := c.Bind(&dataReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data, please check your request body"))
	}
	dataCore := RequestToCore(dataReq)
	data, err := handler.revisionService.Update(idConv, idToken, dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", CoreToResponse(data)))
}

func (handler *RevisionHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	var idConv int
	if id != "" {
		var errConv error
		idConv, errConv = strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("param id is required and must be number"))
		}
	}
	// TODO : get ID from token
	idToken,role := middlewares.ExtractTokenInfo(c, "admin")
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	err := handler.revisionService.Delete(idConv, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
}
