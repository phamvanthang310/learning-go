package dataaccess

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/data-access/dto"
)

type teacherDA struct {
	dbc *bun.DB
}

func (t *teacherDA) Delete(c context.Context, id string) (sql.Result, error) {
	return t.dbc.NewDelete().Model((*dto.Teacher)(nil)).Where("id = ?", id).Exec(c)
}

func (t *teacherDA) GetAll(c context.Context) ([]dto.Teacher, error) {
	var teachers []dto.Teacher
	err := t.dbc.NewSelect().Model(&teachers).Scan(c)

	return teachers, err
}

func (t *teacherDA) GetByUserName(c context.Context, username string) (dto.Teacher, error) {
	teacher := new(dto.Teacher)
	err := t.dbc.NewSelect().Model(teacher).Where("username = ?", username).Scan(c)

	return *teacher, err
}

func (t *teacherDA) Create(ctx context.Context, teacher *dto.Teacher) error {
	_, err := t.dbc.NewInsert().Model(teacher).Exec(ctx)

	return err
}

func NewTeacherDA(dbc *bun.DB) interfaces.TeacherDA {
	return &teacherDA{
		dbc,
	}
}
