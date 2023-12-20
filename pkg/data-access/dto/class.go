package dto

import (
	"github.com/uptrace/bun"
	"time"
)

type (
	Class struct {
		bun.BaseModel `bun:"table:class"`
		ID            string    `bun:"id,pk,autoincrement"`
		Name          string    `bun:"name"`
		StartDate     time.Time `bun:"start_date"`
		EndDate       time.Time `bun:"end_date"`
		CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
		ManagedBy     string    `bun:"managed_by"`
		Teacher       *Teacher  `bun:"rel:belongs-to,join:managed_by=id"`
		Students      []Student `bun:"m2m:student_class,join:Class=Student"`
	}

	StudentClass struct {
		bun.BaseModel `bun:"table:student_class"`
		ClassID       string   `bun:"class_id,pk"`
		Class         *Class   `bun:"rel:belongs-to,join:class_id=id"`
		StudentID     string   `bun:"student_id,pk"`
		Student       *Student `bun:"rel:belongs-to,join:student_id=id"`
	}
)
