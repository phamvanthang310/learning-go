package rest

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"student-service/pkg/application/interfaces"
)

type studentAPI struct {
	studentService interfaces.StudentService
}

func (api *studentAPI) Login(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
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
