package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"student-service/pkg/application/model"
)

type StudentAPI interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
	List(context.Context) ([]model.Student, error)
}
