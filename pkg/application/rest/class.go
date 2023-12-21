package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/utils"
)

type classApi struct {
	service interfaces.ClassService
}

func (c classApi) GetAll(e echo.Context) error {
	claims, err := utils.GetTokenClaims(e)
	if err != nil {
		return echo.ErrForbidden
	}
	result, err := c.service.GetAll(e, claims.ID)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, result)
}

func (c classApi) Create(e echo.Context) error {
	class := new(model.Class)
	if err := utils.BindAndValidate(e, class); err != nil {
		return err
	}

	err := c.service.Create(e, class)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not create new class")
	}

	return e.JSON(http.StatusOK, class)
}

func NewClassApi(s interfaces.ClassService) interfaces.ClassApi {
	return &classApi{s}
}
