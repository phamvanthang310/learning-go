package rest

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/utils"
)

type teacherApi struct {
	service interfaces.TeacherService
}

func (t teacherApi) Login(e echo.Context) error {
	credential := new(model.LoginCredential)
	if err := utils.BindAndValidate(e, credential); err != nil {
		return err
	}

	teacher, err := t.service.FindByUserName(e, credential.Username)
	log.Print(teacher)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credential")
	}

	if err = utils.ComparePassword(teacher.Password, credential.Password); err == nil {
		token := utils.GenerateToken(teacher.User)
		return e.JSON(http.StatusOK, map[string]string{"token": token})
	}
	log.Print(err)

	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credential")
}

func (t teacherApi) Create(e echo.Context) error {
	info := new(model.RegisterInfo)
	if err := utils.BindAndValidate(e, info); err != nil {
		return err
	}

	teacher, err := t.service.Create(e, info)

	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, teacher)
}

func NewTeacherApi(s interfaces.TeacherService) interfaces.TeacherApi {
	return &teacherApi{s}
}
