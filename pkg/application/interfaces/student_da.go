package interfaces

import (
	"context"
	"database/sql"
	"student-service/pkg/data-access/dto"
)

type StudentDA interface {
	Create(ctx context.Context, student *dto.Student) (sql.Result, error)
	FindByUsername(ctx context.Context, username string) (dto.Student, error)
	GetStudents(context.Context) ([]dto.Student, error)
}
