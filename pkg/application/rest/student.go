package rest

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/utils"
)

type studentAPI struct {
	studentService interfaces.StudentService
}

func (api *studentAPI) GetClasses(e echo.Context) error {
	claims, _ := utils.GetTokenClaims(e)
	classes, err := api.studentService.GetClasses(e.Request().Context(), claims.Username)

	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, classes)
}

func (api *studentAPI) Create(e echo.Context) error {
	registerInfo := new(model.RegisterInfo)
	if err := utils.BindAndValidate(e, registerInfo); err != nil {
		return err
	}

	student, err := api.studentService.Create(e.Request().Context(), *registerInfo)
	if err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not register new student")
	}

	return e.JSON(http.StatusCreated, student)
}

func (api *studentAPI) Profile(e echo.Context) error {
	claims, err := utils.GetTokenClaims(e)

	if err != nil {
		return err
	}

	student, err := api.studentService.FindByUsername(e.Request().Context(), claims.Username)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Fail to parse into from token")
	}

	return e.JSON(http.StatusOK, student)
}

func (api *studentAPI) Login(e echo.Context) error {
	credential := new(model.LoginCredential)
	if err := utils.BindAndValidate(e, credential); err != nil {
		return err
	}

	student, err := api.studentService.FindByUsername(e.Request().Context(), credential.Username)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credential")
	}

	if err := utils.ComparePassword(student.Password, credential.Password); err == nil {
		token := utils.GenerateToken(student.User)
		return e.JSON(http.StatusOK, map[string]string{"token": token})
	}

	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credential")
}

func (api *studentAPI) List(c echo.Context) error {
	students, err := api.studentService.GetStudents(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, students)
}

func NewStudentAPI(s interfaces.StudentService) interfaces.StudentAPI {
	return &studentAPI{s}
}
