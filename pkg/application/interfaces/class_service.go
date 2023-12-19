package interfaces

import (
	"github.com/labstack/echo/v4"
	"student-service/pkg/application/model"
)

type ClassService interface {
	Create(ctx echo.Context, c *model.NewClass) (model.Class, error)
}
