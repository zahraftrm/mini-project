package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/zahraftrm/mini-project/app/middlewares"
	training "github.com/zahraftrm/mini-project/features/teacher_training"
	"github.com/zahraftrm/mini-project/helper"

	"github.com/labstack/echo/v4"
)

type TrainingHandler struct {
	trainingService training.TeacherTrainingServiceInterface
}

func New(service training.TeacherTrainingServiceInterface) *TrainingHandler {
	return &TrainingHandler{
		trainingService: service,
	}
}

func (handler *TrainingHandler) GetAll(c echo.Context) error {
	idToken, role := middlewares.ExtractTokenInfo(c,"teacher")
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "teacher" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	// memanggil func di repositories
	results, err := handler.trainingService.GetAll(idToken)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	var trainingResponse []TrainingResponse = CoreToResponseList(results)

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", trainingResponse))
}

func (handler *TrainingHandler) Create(c echo.Context) error {
	idToken, role := middlewares.ExtractTokenInfo(c, "teacher")
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "teacher" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	trainingInput := TrainingRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&trainingInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	//formatting date
	layout_date := "2006-01-02T15:04"
	startDateFormat, errDate := time.Parse(layout_date, fmt.Sprintf("%sT00:00", trainingInput.StartDate))
	if errDate != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed format start date"))
	}
	endDateFormat, errDate := time.Parse(layout_date, fmt.Sprintf("%sT00:00", trainingInput.EndDate))
	if errDate != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed format end date"))
	}
	
	trainingCore := RequestToCore(trainingInput)
	trainingCore.StartDate = startDateFormat
	trainingCore.EndDate = endDateFormat
	trainingCore.TeacherID = uint(idToken)
	err := handler.trainingService.Create(trainingCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
}

func (handler *TrainingHandler) GetById(c echo.Context) error {
	idToken, role := middlewares.ExtractTokenInfo(c, "teacher")
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "teacher" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	idProduct := c.Param("id")
	idProductConv, errConv := strconv.Atoi(idProduct)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed convert id product"))
	}
	// memanggil func di repositories
	result, err := handler.trainingService.GetById(idProductConv, idToken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error read data"))
	}

	var trainingResponse = CoreToResponse(result)

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read training", trainingResponse))
}

func (handler *TrainingHandler) Update(c echo.Context) error {
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
	idToken, role := middlewares.ExtractTokenInfo(c, "teacher")
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "teacher" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	trainingReq := TrainingRequest{}
	errBind := c.Bind(&trainingReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data, please check your request body"))
	}
	trainingCore := RequestToCore(trainingReq)
	data, err := handler.trainingService.Update(idConv, idToken, trainingCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", CoreToResponse(data)))
}

func (handler *TrainingHandler) Delete(c echo.Context) error {
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
	idToken, role := middlewares.ExtractTokenInfo(c, "teacher")
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "teacher" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	err := handler.trainingService.Delete(idConv, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
}
