package service

import (
	"github.com/labstack/echo/v4"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/data-access/dto"
	"student-service/pkg/utils"
)

type teacherService struct {
	db interfaces.TeacherDA
}

func (t teacherService) FindByUserName(e echo.Context, username string) (model.Teacher, error) {
	teacher, err := t.db.GetByUserName(e.Request().Context(), username)
	return mapToTeacherModel(&teacher), err
}

func (t teacherService) Create(e echo.Context, registerInfo *model.RegisterInfo) (*model.Teacher, error) {
	teacherDto := &dto.Teacher{
		Name:     registerInfo.Name,
		Username: registerInfo.Username,
		Password: utils.HashPassword(registerInfo.Password),
		Role:     "teacher",
	}

	if err := t.db.Create(e.Request().Context(), teacherDto); err != nil {
		return nil, err
	}

	teacher := mapToTeacherModel(teacherDto)
	return &teacher, nil
}

func mapToTeacherModel(t *dto.Teacher) model.Teacher {
	return model.Teacher{
		User: model.User{
			ID:        t.ID,
			Name:      t.Name,
			Username:  t.Username,
			Password:  t.Password,
			CreatedAt: t.CreatedAt,
			Role:      t.Role,
		},
	}
}

func NewTeacherService(db interfaces.TeacherDA) interfaces.TeacherService {
	return &teacherService{db}
}
