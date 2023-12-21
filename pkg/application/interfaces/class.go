package interfaces

import (
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
	"student-service/pkg/application/model"
	"student-service/pkg/data-access/dto"
)

type (
	ClassApi interface {
		Create(e echo.Context) error
		GetAll(e echo.Context) error
		DeleteById(e echo.Context) error
	}

	ClassDA interface {
		Create(ctx context.Context, class *dto.Class) (sql.Result, error)
		GetById(ctx context.Context, id string) (*dto.Class, error)
		GetAllManaged(ctx context.Context, username string) ([]dto.Class, error)
		DeleteById(ctx context.Context, id string, userId string) (sql.Result, error)
	}

	ClassService interface {
		Create(ctx echo.Context, c *model.Class) error
		GetAllManaged(ctx echo.Context, username string) ([]model.Class, error)
		DeleteById(context2 echo.Context, id string) (sql.Result, error)
	}
)
