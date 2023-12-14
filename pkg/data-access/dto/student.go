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
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
}
