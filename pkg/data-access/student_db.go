package dataaccess

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"student-service/pkg/data-access/dto"
)

type studentDA struct {
	dbc *bun.DB
}

func (s *studentDA) GetStudents(c context.Context) ([]dto.Student, error) {
	var list []dto.Student
	err := s.dbc.NewSelect().Model(&list).Scan(c)

	return list, err
}

func (s studentDA) Create(ctx context.Context, student *dto.Student) (sql.Result, error) {
	return s.dbc.NewInsert().Model(student).Exec(ctx)
}

func (s studentDA) FindByUsername(ctx context.Context, username string) (dto.Student, error) {
	user := new(dto.Student)
	err := s.dbc.NewSelect().Model(user).Where("username = ?", username).Scan(ctx)
	return *user, err
}

// NewStudentDA creates a new Student Data Access
func NewStudentDA(dbc *bun.DB) *studentDA {
	return &studentDA{dbc}
}
