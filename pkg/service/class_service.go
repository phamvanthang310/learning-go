package service

import (
	"github.com/labstack/echo/v4"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/data-access/dto"
	"student-service/pkg/utils"
)

type classService struct {
	da interfaces.ClassDA
}

func (s *classService) Create(e echo.Context, c *model.Class) error {
	claims, _ := utils.GetTokenClaims(e)
	class := dto.Class{
		Name:      c.Name,
		StartDate: c.StartDate,
		EndDate:   c.EndDate,
		ManagedBy: claims.ID,
	}
	_, err := s.da.Create(e.Request().Context(), &class)

	*c = mapToClassModel(class)
	return err
}

func (s *classService) GetAll(ctx echo.Context) ([]model.Class, error) {
	classDtos, err := s.da.GetAll(ctx.Request().Context())
	var result []model.Class
	for _, item := range classDtos {
		result = append(result, mapToClassModel(item))
	}

	return result, err
}

func mapToClassModel(dto dto.Class) model.Class {
	var students []model.Student
	for _, item := range dto.Students {
		students = append(students, mapToStudentModel(item))
	}
	var teacher *model.Teacher

	if dto.Teacher != nil {
		value := mapToTeacherModel(dto.Teacher)
		teacher = &value
	} else {
		teacher = nil
	}

	return model.Class{
		ID:        dto.ID,
		Name:      dto.Name,
		StartDate: dto.StartDate,
		EndDate:   dto.EndDate,
		CreatedAt: dto.CreatedAt,
		Students:  students,
		ManagedBy: teacher,
	}
}

func NewClassService(da interfaces.ClassDA) interfaces.ClassService {
	return &classService{da}
}
