package dataaccess

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"student-service/pkg/application/interfaces"
	"student-service/pkg/data-access/dto"
)

type classDA struct {
	db *bun.DB
}

func (c *classDA) AssignStudent(ctx context.Context, studentClass []dto.StudentClass) (sql.Result, error) {
	return c.db.NewInsert().Model(&studentClass).Exec(ctx)
}

func (c *classDA) DeleteById(ctx context.Context, id string, userId string) (sql.Result, error) {
	return c.db.NewDelete().
		Model((*dto.Class)(nil)).
		Where("id = ?", id).
		Where("managed_by = ?", userId).Exec(ctx)
}

func (c *classDA) Create(ctx context.Context, class *dto.Class) (sql.Result, error) {
	return c.db.NewInsert().Model(class).Exec(ctx)
}

func (c classDA) GetById(ctx context.Context, id string, userId string) (*dto.Class, error) {
	class := new(dto.Class)
	count, err := c.db.NewSelect().Model(class).
		Where("id = ?", id).
		Where("managed_by = ?", userId).
		Relation("Students").
		ScanAndCount(ctx)

	if count == 0 {
		return nil, err
	}

	return class, err
}

func (c *classDA) GetAllManaged(ctx context.Context, username string) ([]dto.Class, error) {
	var result []dto.Class
	err := c.db.NewSelect().Model(&result).
		Relation("Students").
		Column("class.*").
		Where("class.managed_by = ?", username).
		Scan(ctx)

	return result, err
}

func NewClassDA(db *bun.DB) interfaces.ClassDA {
	return &classDA{db}
}
