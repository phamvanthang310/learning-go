package interfaces

import (
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
	"student-service/pkg/application/model"
	"student-service/pkg/data-access/dto"
)

type (
	TeacherApi interface {
		Create(e echo.Context) error
		Login(e echo.Context) error
		GetAll(e echo.Context) error
		Delete(e echo.Context) error
	}

	TeacherDA interface {
		Create(c context.Context, teacher *dto.Teacher) error
		GetByUserName(c context.Context, username string) (dto.Teacher, error)
		GetAll(c context.Context) ([]dto.Teacher, error)
		Delete(c context.Context, id string) (sql.Result, error)
	}

	TeacherService interface {
		Create(e echo.Context, registerInfo *model.RegisterInfo) (*model.Teacher, error)
		FindByUserName(e echo.Context, username string) (model.Teacher, error)
		GetAll(e echo.Context) ([]model.Teacher, error)
		Delete(e echo.Context, id string) error
	}
)
