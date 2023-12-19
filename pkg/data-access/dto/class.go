package dto

import (
	"github.com/uptrace/bun"
	"time"
)

type (
	Class struct {
		bun.BaseModel `bun:"table:class"`
		ID            int64      `bun:"id,pk,autoincrement"`
		Name          string     `bun:"name"`
		StartDate     time.Time  `bun:"start_date"`
		EndDate       time.Time  `bun:"end_date"`
		CreatedAt     time.Time  `bun:"created_at,default:current_timestamp"`
		Subjects      []*Subject `bun:"m2m:class_subject,join:Class=Subject"`
		ManagedBy     *Teacher   `bun:"rel:has-one,join:managed_by=id"`
	}

	SubjectOfClass struct {
		bun.BaseModel `bun:"table:class_subject"`
		ClassID       int64    `bun:",pk"`
		Class         *Class   `bun:"rel:belongs-to,join:class_id=id"`
		SubjectID     int64    `bun:",pk"`
		Subject       *Subject `bun:"rel:belongs-to,join:subject_id=id"`
	}
)
