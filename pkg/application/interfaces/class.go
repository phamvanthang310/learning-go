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
		GetById(e echo.Context) error
		DeleteById(e echo.Context) error
		AssignStudent(e echo.Context) error
	}

	ClassDA interface {
		Create(ctx context.Context, class *dto.Class) (sql.Result, error)
		GetById(ctx context.Context, id string, userId string) (*dto.Class, error)
		GetAllManaged(ctx context.Context, userId string) ([]dto.Class, error)
		DeleteById(ctx context.Context, id string, userId string) (sql.Result, error)
		AssignStudent(ctx context.Context, students []dto.StudentClass) (sql.Result, error)
	}

	ClassService interface {
		Create(ctx echo.Context, c *model.Class) error
		GetById(ctx echo.Context, id string, userId string) (*model.Class, error)
		GetAllManaged(ctx echo.Context, username string) ([]model.Class, error)
		DeleteById(ctx echo.Context, id string) (sql.Result, error)
		AssignStudent(ctx echo.Context, classId string, studentIds []string, userId string) error
	}
)
