package dataaccess

import (
	"context"
	"database/sql"
	"student-service/pkg/data-access/dto"
	"student-service/pkg/utils"
	"time"

	"github.com/uptrace/bun"
)

type studentDA struct {
	dbc *bun.DB
}

func (s *studentDA) GetStudents(c context.Context) ([]dto.Student, error) {
	// See more: https://bun.uptrace.dev/guide/
	// var list []dto.Student
	// s.dbc.NewSelect().Model(&list).Scan(c)

	return []dto.Student{
		{
			ID:        "123",
			Name:      utils.SampleUtil(),
			CreatedAt: time.Now(),
		},
	}, nil
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
