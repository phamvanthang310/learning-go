package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"student-service/pkg/application/rest"
	"student-service/pkg/config"
	"student-service/pkg/constant"
	dataaccess "student-service/pkg/data-access"
	appMiddlewares "student-service/pkg/middleware"
	"student-service/pkg/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load config
	if _, err := config.LoadConfig(); err != nil {
		panic("Could not load app config")
	}

	sqlDB := dataaccess.InitializeSequelDB()

	// create student data access
	studentDA := dataaccess.NewStudentDA(sqlDB)
	classDA := dataaccess.NewClassDA(sqlDB)
	teacherDA := dataaccess.NewTeacherDA(sqlDB)

	// create student service
	studentService := service.NewStudentService(studentDA)
	classService := service.NewClassService(classDA)
	teacherService := service.NewTeacherService(teacherDA)

	// create APIs
	studentApi := rest.NewStudentAPI(studentService)
	classApi := rest.NewClassApi(classService)
	teacherApi := rest.NewTeacherApi(teacherService)

	server := initializeHTTPServer()

	// Error handler
	server.HTTPErrorHandler = appMiddlewares.ErrorHandler
	server.Validator = &appMiddlewares.CustomValidator{
		Validator: validator.New(),
	}

	// Index page
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!")
	})

	// Public routes
	server.POST("/login", teacherApi.Login)
	server.POST("/student/login", studentApi.Login)

	// student
	authenticated := server.Group("/student", appMiddlewares.Authentication)
	authenticated.GET("/profile", studentApi.Profile)
	authenticated.GET("/class", studentApi.GetClasses)

	// teacher
	teacherRoute := server.Group("/teacher", appMiddlewares.Authentication,
		appMiddlewares.Authorization(constant.TeacherRole, constant.AdminRole))
	teacherRoute.POST("/class", classApi.Create)
	teacherRoute.GET("/class", classApi.GetAll)
	teacherRoute.DELETE("/class/:id", classApi.DeleteById)
	teacherRoute.GET("/class/:id", classApi.GetById)
	teacherRoute.POST("/class/assign/:id", classApi.AssignStudent)

	// admin
	adminRoute := server.Group("admin", appMiddlewares.Authentication,
		appMiddlewares.Authorization(constant.AdminRole))
	adminRoute.POST("/student", studentApi.Create)
	adminRoute.POST("/teacher", teacherApi.Create)
	adminRoute.GET("/students", studentApi.List)

	// Start listening
	server.Logger.Fatal(server.Start("127.0.0.1:8080"))
}

func initializeHTTPServer() *echo.Echo {
	// Echo instance customization
	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.Logger())  // Logger middleware
	e.Use(middleware.Recover()) // Panic recover middleware

	return e
}
