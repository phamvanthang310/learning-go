package interfaces

import (
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
	"student-service/pkg/application/model"
	"student-service/pkg/data-access/dto"
)

type (
	StudentAPI interface {
		Login(ctx echo.Context) error
		List(ctx echo.Context) error
	}

	StudentDA interface {
		Create(ctx context.Context, student *dto.Student) (sql.Result, error)
		FindByUsername(ctx context.Context, username string) (dto.Student, error)
		GetStudents(context.Context) ([]dto.Student, error)
	}

	StudentService interface {
		Create(ctx context.Context, info model.RegisterInfo) (model.Student, error)
		FindByUsername(ctx context.Context, username string) (model.Student, error)
		GetStudents(context.Context) ([]model.Student, error)
	}
)
