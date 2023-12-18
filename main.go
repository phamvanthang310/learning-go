package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"student-service/pkg/application/rest"
	"student-service/pkg/config"
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

	// create student service
	studentService := service.NewStudentService(studentDA)

	// create APIs
	studentAPI := rest.NewStudentAPI(studentService)
	authApi := rest.NewAuthApi(studentService)

	server := initializeHTTPServer()

	// Index page
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!")
	})

	// Routes
	server.POST("/register", authApi.Register)
	server.POST("/login", authApi.Login)

	// Error handler
	server.HTTPErrorHandler = appMiddlewares.ErrorHandler
	server.Validator = &appMiddlewares.CustomValidator{
		Validator: validator.New(),
	}

	// authenticated endpoints
	authenticated := server.Group("", appMiddlewares.Auth)
	authenticated.GET("/info", authApi.Info)
	authenticated.GET("/students", studentAPI.List)

	// teacher
	// authenticated.GET("/classes", example.Handle)
	// authenticated.POST("/classes", example.Handle)
	// authenticated.PATCH("/classes/:classId", example.Handle)

	// admin
	// authenticated.GET("/users", example.Handle)
	// authenticated.POST("/users", example.Handle)
	// authenticated.PATCH("/users/:userId", example.Handle)

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
