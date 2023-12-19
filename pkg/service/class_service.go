package service

import (
	"github.com/labstack/echo/v4"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/data-access/dto"
)

type classService struct {
	da interfaces.ClassDA
}

func (s *classService) Create(ctx echo.Context, c *model.NewClass) (model.Class, error) {
	class := dto.Class{
		Name:      c.Name,
		StartDate: c.StartDate,
		EndDate:   c.EndDate,
	}

	_, err := s.da.Create(ctx.Request().Context(), &class)

	return mapToClassModel(class), err
}

func mapToClassModel(dto dto.Class) model.Class {
	return model.Class{
		ID:        dto.ID,
		Name:      dto.Name,
		StartDate: dto.StartDate,
		EndDate:   dto.EndDate,
		CreatedAt: dto.CreatedAt,
	}
}

func NewClassService(da interfaces.ClassDA) interfaces.ClassService {
	return &classService{da}
}
