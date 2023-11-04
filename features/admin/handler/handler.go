package handler

import (
	"net/http"
	"strings"

	"github.com/zahraftrm/mini-project/app/middlewares"
	"github.com/zahraftrm/mini-project/constants"
	"github.com/zahraftrm/mini-project/features/admin"
	"github.com/zahraftrm/mini-project/helper"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminService admin.AdminServiceInterface
}

func New(service admin.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{
		adminService: service,
	}
}

func (handler *AdminHandler) GetAllAdmin(c echo.Context) error {
	_, role := middlewares.ExtractTokenUserId(c)
	if role != constants.RolesAdmin {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	
	// memanggil func di repositories
	results, err := handler.adminService.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	var adminsResponse []AdminResponse
	for _, value := range results {
		adminsResponse = append(adminsResponse, AdminResponse{
			Id:        value.Id,
			Name:      value.Name,
			Email:     value.Email,
			CreatedAt: value.CreatedAt,
		})
	}

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data", adminsResponse))
}

func (handler *AdminHandler) CreateAdmin(c echo.Context) error {
	adminInput := AdminRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&adminInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	// mapping dari request ke core
	adminCore := admin.Core{
		Name:     adminInput.Name,
		Email:    adminInput.Email,
		Password: adminInput.Password,
	}
	
	err := handler.adminService.Create(adminCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
}

func (handler *AdminHandler) Login(c echo.Context) error {
	loginInput := AuthRequest{}
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	dataLogin, token, err := handler.adminService.Login(loginInput.Email, loginInput.Password)
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

func (handler *AdminHandler) GetProfile(c echo.Context) error {
	idToken, role := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != constants.RolesAdmin {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	// memanggil func di repositories
	result, err := handler.adminService.GetById(idToken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error read data"))
	}

	var adminsResponse = CoreToResponse(result)

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read profile", adminsResponse))
}

func (handler *AdminHandler) Update(c echo.Context) error {
	idToken, role := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != constants.RolesAdmin {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	adminReq := AdminRequest{}
	errBind := c.Bind(&adminReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data, please check your request body"))
	}
	adminCore := RequestToCore(adminReq)
	data, err := handler.adminService.Update(idToken, adminCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", CoreToResponse(data)))
}

func (handler *AdminHandler) Delete(c echo.Context) error {
	// TODO : get ID from token
	idToken, role := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}
	if role != constants.RolesAdmin {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("unauthorized"))
	}

	err := handler.adminService.Delete(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success deactivate account"))
}
