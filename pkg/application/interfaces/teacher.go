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
		Login(e echo.Context) error
		GetAll(c echo.Context) error
	}

	TeacherDA interface {
		Create(ctx context.Context, teacher *dto.Teacher) error
		GetByUserName(c context.Context, username string) (dto.Teacher, error)
		GetAll(c context.Context) ([]dto.Teacher, error)
	}

	TeacherService interface {
		Create(e echo.Context, registerInfo *model.RegisterInfo) (*model.Teacher, error)
		FindByUserName(e echo.Context, username string) (model.Teacher, error)
		GetAll(c echo.Context) ([]model.Teacher, error)
	}
)
