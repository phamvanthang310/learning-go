package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"student-service/pkg/application/model"
	"student-service/pkg/data-access/dto"
)

type (
	TeacherApi interface {
		Create(e echo.Context) error
	}

	TeacherDA interface {
		Create(ctx context.Context, teacher *dto.Teacher) error
		GetByUserName(c context.Context, username string) (dto.Teacher, error)
		GetAll(c context.Context) ([]dto.Teacher, error)
	}

	TeacherService interface {
		Create(e echo.Context, teacher *model.Teacher) error
	}
)
