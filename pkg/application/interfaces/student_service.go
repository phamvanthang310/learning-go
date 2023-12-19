package interfaces

import (
	"context"
	"student-service/pkg/application/model"
)

type StudentService interface {
	Create(ctx context.Context, info model.RegisterInfo) (model.Student, error)
	FindByUsername(ctx context.Context, username string) (model.Student, error)
	GetStudents(context.Context) ([]model.Student, error)
}
