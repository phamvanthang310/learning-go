package interfaces

import (
	"context"
	"database/sql"
	"student-service/pkg/data-access/dto"
)

type ClassDA interface {
	Create(ctx context.Context, class *dto.Class) (sql.Result, error)
	GetById(ctx context.Context, id string) (*dto.Class, error)
}
