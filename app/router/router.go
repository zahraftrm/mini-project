package router

import (
	_adminData "github.com/zahraftrm/mini-project/features/admin/data"
	_adminHandler "github.com/zahraftrm/mini-project/features/admin/handler"
	_adminService "github.com/zahraftrm/mini-project/features/admin/service"

	_teacherData "github.com/zahraftrm/mini-project/features/teacher/data"
	_teacherHandler "github.com/zahraftrm/mini-project/features/teacher/handler"
	_teacherService "github.com/zahraftrm/mini-project/features/teacher/service"

	// _trainingData "github.com/zahraftrm/mini-training/features/teacher_training/data"
	// _trainingHandler "github.com/zahraftrm/mini-training/features/teacher_training/handler"
	// _trainingService "github.com/zahraftrm/mini-training/features/teacher_training/service"

	_revisionData "github.com/zahraftrm/mini-project/features/revision/data"
	_revisionHandler "github.com/zahraftrm/mini-project/features/revision/handler"
	_revisionService "github.com/zahraftrm/mini-project/features/revision/service"

	"github.com/zahraftrm/mini-project/app/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	adminData := _adminData.New(db)
	adminService := _adminService.New(adminData)
	adminHandlerAPI := _adminHandler.New(adminService)

	teacherData := _teacherData.New(db)
	teacherService := _teacherService.New(teacherData)
	teacherHandlerAPI := _teacherHandler.New(teacherService)

	// trainingData := _trainingData.New(db)
	// trainingService := _trainingService.New(trainingData)
	// trainingHandlerAPI := _trainingHandler.New(trainingService)

	revisionData := _revisionData.New(db)
	revisionService := _revisionService.New(revisionData)
	revisionHandlerAPI := _revisionHandler.New(revisionService)

	//e.GET("/admins", userHandlerAPI.GetAllAdmins, middlewares.JWTMiddlewareWithRole(1))

	// endpoint admins accessed by admins
	e.GET("/admins", adminHandlerAPI.GetAllAdmin, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("admin"))
	e.GET("/admins/profile", adminHandlerAPI.GetProfile, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("admin"))
	e.PUT("/admins", adminHandlerAPI.Update, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("admin"))
	e.DELETE("/admins", adminHandlerAPI.Delete, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("admin"))
	e.POST("/admins", adminHandlerAPI.CreateAdmin)
	e.POST("/admins/login", adminHandlerAPI.Login)

	// endpoint teachers accessed by teacher
	e.GET("/teachers/profile", teacherHandlerAPI.GetProfile, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("teacher"))
	e.PUT("/teachers", teacherHandlerAPI.Update, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("teacher"))
	e.POST("/teachers/login", teacherHandlerAPI.Login)

	// endpoint teachers accessed by admin
	e.GET("/admins/teachers", teacherHandlerAPI.GetAllTeacher, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("admin"))
	e.GET("/admins/teachers/:id", teacherHandlerAPI.GetProfile, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("admin"))
	e.POST("/admins/teachers", teacherHandlerAPI.CreateTeacher, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("admin"))
	e.PUT("/admins/teachers/:id", teacherHandlerAPI.Update, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("admin"))
	e.DELETE("/admins/teachers/:id", teacherHandlerAPI.Delete, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("admin"))

	// // e.GET("/admins/:adminid/trainings", trainingHandlerAPI.GetProjectByAdminId, middlewares.JWTMiddleware())
	// e.POST("/trainings", trainingHandlerAPI.Create, middlewares.JWTMiddleware())
	// e.GET("/trainings", trainingHandlerAPI.GetAll, middlewares.JWTMiddleware())
	// e.GET("/trainings/:id", trainingHandlerAPI.GetById, middlewares.JWTMiddleware())

	// endpoint revisions accessed by admin
	e.POST("/admins/revisions", revisionHandlerAPI.Create, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("admin"))
	e.PUT("/admins/revisions/:id", revisionHandlerAPI.Update, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("admin"))
	e.DELETE("/admins/revision/:id", revisionHandlerAPI.Delete, middlewares.JWTMiddleware(), middlewares.RoleAuthorization("admin"))
}