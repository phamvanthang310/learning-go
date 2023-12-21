package service

import (
	"context"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/application/model"
	"student-service/pkg/data-access/dto"
	"student-service/pkg/utils"
)

type studentService struct {
	db interfaces.StudentDA
}

func (s *studentService) GetClasses(ctx context.Context, username string) ([]model.Class, error) {
	classDtos, err := s.db.GetClasses(ctx, username)
	result := make([]model.Class, len(classDtos))

	for i, v := range classDtos {
		result[i] = mapToClassModel(v)
	}

	return result, err
}

func (s *studentService) GetStudents(c context.Context) ([]model.Student, error) {
	list, err := s.db.GetStudents(c)
	if err != nil {
		return nil, err
	}

	// map from dto to model
	result := make([]model.Student, len(list))
	for i, v := range list {
		result[i] = mapToStudentModel(v)
	}
	return result, nil
}

func (s *studentService) Create(ctx context.Context, info model.RegisterInfo) (model.Student, error) {
	student := dto.Student{
		Name:     info.Name,
		Username: info.Username,
		Password: utils.HashPassword(info.Password),
	}

	_, err := s.db.Create(ctx, &student)
	return mapToStudentModel(student), err
}

func (s *studentService) FindByUsername(ctx context.Context, username string) (model.Student, error) {
	student, err := s.db.FindByUsername(ctx, username)

	return mapToStudentModel(student), err
}

func mapToStudentModel(s dto.Student) model.Student {
	classes := make([]model.Class, len(s.Classes))

	for i, v := range s.Classes {
		classes[i] = mapToClassModel(v)
	}

	return model.Student{
		User: model.User{
			ID:        s.ID,
			Name:      s.Name,
			Username:  s.Username,
			Password:  s.Password,
			CreatedAt: s.CreatedAt,
		},
		Classes: classes,
	}
}

func NewStudentService(db interfaces.StudentDA) interfaces.StudentService {
	return &studentService{db}
}
