package service

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/data-access/dto"
	"student-service/pkg/utils"
)

type classService struct {
	da interfaces.ClassDA
}

func (s *classService) GetById(e echo.Context, id string, userId string) (*model.Class, error) {
	classDto, err := s.da.GetById(e.Request().Context(), id, userId)

	if classDto == nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Class not found")
	}

	result := mapToClassModel(*classDto)
	return &result, err
}

func (s *classService) AssignStudent(e echo.Context, classId string, studentIds []string, userId string) error {
	_, err := s.GetById(e, classId, userId)
	if err != nil {
		return err
	}

	studentClass := make([]dto.StudentClass, len(studentIds))
	for i, v := range studentIds {
		studentClass[i] = dto.StudentClass{
			StudentID: v,
			ClassID:   classId,
		}
	}

	_, err = s.da.AssignStudent(e.Request().Context(), studentClass)
	if err != nil {
		log.Print(err)
		return echo.ErrInternalServerError
	}

	return nil
}

func (s *classService) DeleteById(e echo.Context, id string) (sql.Result, error) {
	claims, _ := utils.GetTokenClaims(e)
	return s.da.DeleteById(e.Request().Context(), id, claims.ID)
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

func (s *classService) GetAllManaged(ctx echo.Context, userId string) ([]model.Class, error) {
	classDtos, err := s.da.GetAllManaged(ctx.Request().Context(), userId)
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
