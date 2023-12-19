package dataaccess

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
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

func NewClassDA(db *bun.DB) *classDA {
	return &classDA{db}
}
