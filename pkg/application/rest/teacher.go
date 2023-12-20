package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/utils"
)

type teacherApi struct {
	service interfaces.TeacherService
}

func (t teacherApi) Create(e echo.Context) error {
	teacher := new(model.Teacher)
	if err := utils.BindAndValidate(e, teacher); err != nil {
		return err
	}

	if err := t.service.Create(e, teacher); err != nil {
		return err
	}

	return e.JSON(http.StatusOK, teacher)
}

func NewTeacherApi(s interfaces.TeacherService) interfaces.TeacherApi {
	return &teacherApi{s}
}
