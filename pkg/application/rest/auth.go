package rest

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
)

type authRestApi struct {
	service interfaces.StudentService
}

func (a authRestApi) Register(e echo.Context) error {
	registerInfo := new(model.RegisterInfo)

	if err := e.Bind(&registerInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON request")
	}

	if err := a.service.Create(e.Request().Context(), *registerInfo); err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not register new student")
	}

	return e.JSON(http.StatusOK, registerInfo)
}

func NewAuthApi(service interfaces.StudentService) *authRestApi {
	return &authRestApi{service}
}
