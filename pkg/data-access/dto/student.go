package dto

import (
	"time"

	"github.com/uptrace/bun"
)

type Student struct {
	bun.BaseModel `bun:"table:student"`
	ID            string    `bun:"id,pk,autoincrement"`
	Name          string    `bun:"name"`
	Username      string    `bun:"username"`
	Password      string    `bun:"password"`
	Role          string    `bun:"role"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
	Classes       []Class   `bun:"m2m:student_class,join:Student=Class"`
}
