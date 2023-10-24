package handler

import (
	"net/http"
	"strings"

	"github.com/zahraftrm/mini-project/app/middlewares"
	"github.com/zahraftrm/mini-project/features/teacher"
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
	teacherInput := TeacherRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&teacherInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	// mapping dari request ke core
	teacherCore := teacher.Core{
		NUPTK:    	teacherInput.NUPTK,
		Name:     	teacherInput.Name,
		Phone:    	teacherInput.Phone,
		Expertise:  teacherInput.Expertise,
		Email:    	teacherInput.Email,
		Password: 	teacherInput.Password,
	}
	err := handler.teacherService.Create(teacherCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
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
	idToken := middlewares.ExtractTokenTeacherId(c)

	// memanggil func di repositories
	result, err := handler.teacherService.GetById(idToken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error read data"))
	}

	var teachersResponse = CoreToResponse(result)

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read profile", teachersResponse))
}

func (handler *TeacherHandler) Update(c echo.Context) error {
	// id := c.Param("id")
	// var idConv int
	// if id != "" {
	// 	var errConv error
	// 	idConv, errConv = strconv.Atoi(id)
	// 	if errConv != nil {
	// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse("param id is required and must be number"))
	// 	}
	// }

	// TODO : get ID from token
	idToken := middlewares.ExtractTokenTeacherId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	teacherReq := TeacherRequest{}
	errBind := c.Bind(&teacherReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data, please check your request body"))
	}
	teacherCore := RequestToCore(teacherReq)
	data, err := handler.teacherService.Update(idToken, teacherCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", CoreToResponse(data)))
}

func (handler *TeacherHandler) Delete(c echo.Context) error {
	// TODO : get ID from token
	idToken := middlewares.ExtractTokenTeacherId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	err := handler.teacherService.Delete(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success deactivate account"))
}
