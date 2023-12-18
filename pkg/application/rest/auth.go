package rest

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/utils"
)

type authRestApi struct {
	service interfaces.StudentService
}

func (a authRestApi) Register(e echo.Context) error {
	registerInfo := new(model.RegisterInfo)
	if err := utils.BindAndValidate(registerInfo, e); err != nil {
		return err
	}

	if err := a.service.Create(e.Request().Context(), *registerInfo); err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not register new student")
	}

	return e.JSON(http.StatusCreated, registerInfo)
}

func (a authRestApi) Login(e echo.Context) error {
	credential := new(model.LoginCredential)
	if err := utils.BindAndValidate(credential, e); err != nil {
		return err
	}

	student, err := a.service.FindByUsername(e.Request().Context(), credential.Username)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Fail to login")
	}

	if err := utils.Compare(student.Password, credential.Password); err == nil {
		token := utils.GenerateToken(student)
		return e.JSON(http.StatusOK, map[string]string{"token": token})
	}

	return echo.NewHTTPError(http.StatusUnauthorized, "Fail to login")
}

func (a authRestApi) Info(e echo.Context) error {
	claims, err := utils.GetTokenClaims(e)

	if err != nil {
		return err
	}

	student, err := a.service.FindByUsername(e.Request().Context(), claims.Username)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Fail to parse into from token")
	}

	return e.JSON(http.StatusOK, student)
}

func NewAuthApi(service interfaces.StudentService) *authRestApi {
	return &authRestApi{service}
}
