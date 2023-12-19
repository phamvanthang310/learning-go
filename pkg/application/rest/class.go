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

func (c classApi) Create(e echo.Context) error {
	newClass := new(model.NewClass)
	if err := utils.BindAndValidate(e, newClass); err != nil {
		return err
	}

	class, err := c.service.Create(e, newClass)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not create new class")
	}

	return e.JSON(http.StatusOK, class)
}

func NewClassApi(s interfaces.ClassService) interfaces.ClassApi {
	return &classApi{s}
}
