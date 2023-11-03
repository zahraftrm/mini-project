package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/zahraftrm/mini-project/app/middlewares"
	training "github.com/zahraftrm/mini-project/features/training"
	"github.com/zahraftrm/mini-project/helper"

	"github.com/labstack/echo/v4"
)

type TrainingHandler struct {
	trainingService training.TrainingServiceInterface
}

func New(service training.TrainingServiceInterface) *TrainingHandler {
	return &TrainingHandler{
		trainingService: service,
	}
}

func (handler *TrainingHandler) Create(c echo.Context) error {
	idToken, role := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "teacher"{
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
	
	trainingCore, err := RequestToCore(trainingInput, role)
	if err != nil {
	    return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	trainingCore.StartDate = startDateFormat
	trainingCore.EndDate = endDateFormat

	trainingCore.TeacherID = uint(idToken)

	err = handler.trainingService.Create(trainingCore, role)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
}

func (handler *TrainingHandler) GetAll(c echo.Context) error {
	idToken, role := middlewares.ExtractTokenUserId(c)
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

func (handler *TrainingHandler) GetAllByAdmin(c echo.Context) error {
	idToken, role := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	// memanggil func di repositories
	results, err := handler.trainingService.GetAllByAdmin()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	var trainingResponse []TrainingResponse = CoreToResponseList(results)

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", trainingResponse))
}

func (handler *TrainingHandler) GetById(c echo.Context) error {
	idToken, role := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "teacher" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	idTraining := c.Param("id")
	idTrainingConv, errConv := strconv.Atoi(idTraining)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed convert id training"))
	}
	// memanggil func di repositories
	result, err := handler.trainingService.GetById(idTrainingConv, idToken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error read data"))
	}

	var trainingResponse = CoreToResponse(result)

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read training", trainingResponse))
}

func (handler *TrainingHandler) GetByIdTeacher(c echo.Context) error {
	idToken, role := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	idTeacher := c.Param("id")
	idTeacherConv, errConv := strconv.Atoi(idTeacher)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed convert id training"))
	}
	// memanggil func di repositories
	result, err := handler.trainingService.GetByIdTeacher(idTeacherConv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error read data"))
	}

	var trainingResponse []TrainingResponse= CoreToResponseList(result)

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read training", trainingResponse))
}

func (handler *TrainingHandler) GetByIdTraining(c echo.Context) error {
	idToken, role := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	idTraining := c.Param("id")
	idTrainingConv, errConv := strconv.Atoi(idTraining)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed convert id training"))
	}
	// memanggil func di repositories
	result, err := handler.trainingService.GetByIdTraining(idTrainingConv)

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
	idToken, role := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	trainingReq := TrainingRequest{}
	errBind := c.Bind(&trainingReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data, please check your request body"))
	}

	trainingCore, err := RequestToCore(trainingReq, role)
	if err != nil {
	    return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	trainingCore.AdminID = uint(idToken)

	data, email, name, err := handler.trainingService.Update(idConv, idToken, role, trainingCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	err = helper.SendEmailTrainingStatus(email, name, trainingReq.Status, trainingReq.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error send email"))
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
	idToken, role := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != "teacher" && role != "admin"{
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	var err error
	if role == "teacher" {
		err = handler.trainingService.Delete(idConv, idToken)
	}else if role == "admin" {
		err = handler.trainingService.DeleteByAdmin(idConv)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
}
