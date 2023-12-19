package dataaccess

import (
	"context"
	"github.com/uptrace/bun"
	"student-service/pkg/data-access/dto"
)

type teacherDA struct {
	dbc *bun.DB
}

func (t *teacherDA) GetTeachers(c context.Context) ([]dto.Teacher, error) {
	var teachers []dto.Teacher
	err := t.dbc.NewSelect().Model(&teachers).Scan(c)

	return teachers, err
}

func (t *teacherDA) GetByUserName(c context.Context, username string) (dto.Teacher, error) {
	teacher := new(dto.Teacher)
	_, err := t.dbc.NewSelect().Model(teacher).Where("id = ?", username).Exec(c)

	return *teacher, err
}

func (t *teacherDA) Create(ctx context.Context, teacher *dto.Teacher) (dto.Teacher, error) {
	_, err := t.dbc.NewInsert().Model(teacher).Exec(ctx)

	return *teacher, err
}

func NewTeacherDA(dbc *bun.DB) *teacherDA {
	return &teacherDA{
		dbc,
	}
}
