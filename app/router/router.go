package router

import (
	_adminData "github.com/zahraftrm/mini-project/features/admin/data"
	_adminHandler "github.com/zahraftrm/mini-project/features/admin/handler"
	_adminService "github.com/zahraftrm/mini-project/features/admin/service"

	_teacherData "github.com/zahraftrm/mini-project/features/teacher/data"
	_teacherHandler "github.com/zahraftrm/mini-project/features/teacher/handler"
	_teacherService "github.com/zahraftrm/mini-project/features/teacher/service"

	_trainingData "github.com/zahraftrm/mini-project/features/training/data"
	_trainingHandler "github.com/zahraftrm/mini-project/features/training/handler"
	_trainingService "github.com/zahraftrm/mini-project/features/training/service"

	_revisionData "github.com/zahraftrm/mini-project/features/revision/data"
	_revisionHandler "github.com/zahraftrm/mini-project/features/revision/handler"
	_revisionService "github.com/zahraftrm/mini-project/features/revision/service"

	// "github.com/zahraftrm/mini-project/features/openAI/handler"
	handlers "github.com/zahraftrm/mini-project/features/recommendation/handler"
	"github.com/zahraftrm/mini-project/features/recommendation/service"

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

	trainingData := _trainingData.New(db)
	trainingService := _trainingService.New(trainingData)
	trainingHandlerAPI := _trainingHandler.New(trainingService)

	revisionData := _revisionData.New(db)
	revisionService := _revisionService.New(revisionData)
	revisionHandlerAPI := _revisionHandler.New(revisionService)

	recommendationService := service.NewRecommendationService()
	recommendationHandler := handlers.NewRecommendationHandler(recommendationService)
	e.POST("/recommendations", recommendationHandler.Recommendation)

	// endpoint admins accessed by admins
	e.GET("/admins", adminHandlerAPI.GetAllAdmin, middlewares.JWTMiddleware())
	e.GET("/admins/profile", adminHandlerAPI.GetProfile, middlewares.JWTMiddleware())
	e.PUT("/admins/profile", adminHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/admins", adminHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.POST("/admins", adminHandlerAPI.CreateAdmin)
	e.POST("/admins/login", adminHandlerAPI.Login)

	// endpoint teachers accessed by teacher
	e.GET("/teachers/profile", teacherHandlerAPI.GetProfile, middlewares.JWTMiddleware())
	e.PUT("/teachers/profile", teacherHandlerAPI.Update, middlewares.JWTMiddleware())
	e.POST("/teachers/login", teacherHandlerAPI.Login)

	// endpoint teachers accessed by admin
	e.GET("/admins/teachers", teacherHandlerAPI.GetAllTeacher, middlewares.JWTMiddleware())
	e.GET("/admins/teachers/:id", teacherHandlerAPI.GetProfile, middlewares.JWTMiddleware())
	e.POST("/admins/teachers", teacherHandlerAPI.CreateTeacher, middlewares.JWTMiddleware())
	e.PUT("/admins/teachers/:id", teacherHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/admins/teachers/:id", teacherHandlerAPI.Delete, middlewares.JWTMiddleware())

	// endpoint trainings accessed by teacher
	e.GET("/teachers/trainings", trainingHandlerAPI.GetAll, middlewares.JWTMiddleware())
	e.GET("/teachers/trainings/:id", trainingHandlerAPI.GetById, middlewares.JWTMiddleware())
	e.POST("/teachers/trainings", trainingHandlerAPI.Create, middlewares.JWTMiddleware())
	e.DELETE("/teachers/trainings/:id", trainingHandlerAPI.Delete, middlewares.JWTMiddleware())

	// endpoint trainings accessed by admin
	e.GET("/admins/trainings", trainingHandlerAPI.GetAllByAdmin, middlewares.JWTMiddleware())
	e.GET("/admins/teachers/:id/trainings", trainingHandlerAPI.GetByIdTeacher, middlewares.JWTMiddleware())
	e.GET("/admins/trainings/:id", trainingHandlerAPI.GetByIdTraining, middlewares.JWTMiddleware())
	e.PUT("/admins/trainings/:id", trainingHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/admins/trainings/:id", trainingHandlerAPI.Delete, middlewares.JWTMiddleware())

	// endpoint revisions accessed by admin
	e.POST("/admins/revisions", revisionHandlerAPI.Create, middlewares.JWTMiddleware())
	e.PUT("/admins/revisions/:id", revisionHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/admins/revisions/:id", revisionHandlerAPI.Delete, middlewares.JWTMiddleware())
}