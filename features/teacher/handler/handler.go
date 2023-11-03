package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/zahraftrm/mini-project/app/middlewares"
	"github.com/zahraftrm/mini-project/constants"
	"github.com/zahraftrm/mini-project/features/teacher"
	"github.com/zahraftrm/mini-project/features/teacher/service"
	"github.com/zahraftrm/mini-project/helper"

	"github.com/labstack/echo/v4"
)

type TeacherHandler struct {
	teacherService teacher.TeacherServiceInterface
}

func New(service teacher.TeacherServiceInterface) *TeacherHandler {
	return &TeacherHandler{
		teacherService: service,
	}
}

func (handler *TeacherHandler) GetAllTeacher(c echo.Context) error {
	_, role := middlewares.ExtractTokenUserId(c)
	if role != constants.RolesAdmin {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	
	// memanggil func di repositories
	results, err := handler.teacherService.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	var teachersResponse []TeacherResponse
	for _, value := range results {
		teachersResponse = append(teachersResponse, TeacherResponse{
			Id:        value.Id,
			NUPTK:	   value.NUPTK,
			Name:      value.Name,
			Phone:     value.Phone,
			Expertise: value.Expertise,
			Email:     value.Email,
			CreatedAt: value.CreatedAt,
		})
	}

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", teachersResponse))
}

func (handler *TeacherHandler) CreateTeacher(c echo.Context) error {
	idToken, role := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != constants.RolesAdmin {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	
	teacherInput := TeacherRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&teacherInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	password := service.RandomPassword()

	// mapping dari request ke core
	teacherCore := teacher.Core{
		NUPTK:    	teacherInput.NUPTK,
		Name:     	teacherInput.Name,
		Phone:    	teacherInput.Phone,
		Expertise:  teacherInput.Expertise,
		Email:    	teacherInput.Email,
		Password: 	password,
	}

	err := handler.teacherService.Create(teacherCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
	}

	err = helper.SendEmailNewTeacherAccount(teacherInput.Email, teacherInput.Name, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error send email"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
}

func (handler *TeacherHandler) Login(c echo.Context) error {
	loginInput := AuthRequest{}
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	dataLogin, token, err := handler.teacherService.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "login failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error login, intrnal server error"))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("login success", map[string]any{
		"token": token,
		"email": dataLogin.Email,
		"id":    dataLogin.Id,
	}))

}

func (handler *TeacherHandler) GetProfile(c echo.Context) error {
	id := c.Param("id")
	var idConv int
	if id != "" {
		var errConv error
		idConv, errConv = strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("param id is required and must be number"))
		}
	}
	
	idToken, role := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != constants.RolesAdmin && role != constants.RolesTeacher {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	// memanggil func di repositories
	var result teacher.Core
	var err error
	if role == constants.RolesTeacher {
		result, err = handler.teacherService.GetById(idToken)
	} else if role == constants.RolesAdmin {
		result, err = handler.teacherService.GetById(idConv)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error read data"))
	}

	var teachersResponse = CoreToResponse(result)

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read profile", teachersResponse))
}

func (handler *TeacherHandler) Update(c echo.Context) error {
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
	if role != constants.RolesAdmin && role != constants.RolesTeacher {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	teacherReq := TeacherRequest{}
	errBind := c.Bind(&teacherReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data, please check your request body"))
	}
	teacherCore := RequestToCore(teacherReq)
	
	var data teacher.Core
	var err error
	if role == constants.RolesTeacher{
		data, err = handler.teacherService.Update(idToken, teacherCore)
	} else if role == constants.RolesAdmin {
		data, err = handler.teacherService.Update(idConv, teacherCore)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", CoreToResponse(data)))
}

func (handler *TeacherHandler) Delete(c echo.Context) error {
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
	if role != constants.RolesAdmin {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	var err error
	if role == constants.RolesTeacher {
		err = handler.teacherService.Delete(idToken)
	} else if role == constants.RolesAdmin {
		err = handler.teacherService.Delete(idConv)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success deactivate account"))
}
