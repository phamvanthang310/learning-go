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

func (c *classDA) Create(ctx context.Context, class *dto.Class) (sql.Result, error) {
	return c.db.NewInsert().Model(class).Exec(ctx)
}

func (c classDA) GetById(ctx context.Context, id string) (*dto.Class, error) {
	class := new(dto.Class)
	err := c.db.NewSelect().Model(class).Where("id = ?", id).Scan(ctx)
	return class, err
}

func (c *classDA) GetAll(ctx context.Context) ([]dto.Class, error) {
	var result []dto.Class
	err := c.db.NewSelect().Model(&result).Relation("Students").Column("class.*").Relation("Teacher").Scan(ctx)
	return result, err
}

func NewClassDA(db *bun.DB) interfaces.ClassDA {
	return &classDA{db}
}
