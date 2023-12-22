package rest

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/utils"
)

type classApi struct {
	service interfaces.ClassService
}

func (c classApi) GetById(e echo.Context) error {
	classId := e.Param("id")

	claims, err := utils.GetTokenClaims(e)
	if err != nil {
		return err
	}

	result, err := c.service.GetById(e, classId, claims.ID)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, result)
}

func (c classApi) AssignStudent(e echo.Context) error {
	classId := e.Param("id")

	claims, err := utils.GetTokenClaims(e)
	if err != nil {
		return err
	}

	assignStudent := new(model.AssignStudent)
	if err := utils.BindAndValidate(e, assignStudent); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.service.AssignStudent(e, classId, assignStudent.StudentIds, claims.ID); err != nil {
		return err
	}

	class, _ := c.service.GetById(e, classId, claims.ID)
	return e.JSON(http.StatusOK, class)
}

func (c classApi) DeleteById(e echo.Context) error {
	id := e.Param("id")

	result, err := c.service.DeleteById(e, id)

	if err != nil {
		log.Print(err)
		return echo.ErrInternalServerError
	}

	if no, err := result.RowsAffected(); err != nil {
		log.Print(err)
		return echo.ErrInternalServerError
	} else if no == 0 {
		return echo.ErrNotFound
	}

	return e.JSON(http.StatusNoContent, "")
}

func (c classApi) GetAll(e echo.Context) error {
	claims, err := utils.GetTokenClaims(e)
	if err != nil {
		return err
	}
	result, err := c.service.GetAllManaged(e, claims.ID)
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
