package dto

import (
	"github.com/uptrace/bun"
	"time"
)

type Subject struct {
	bun.BaseModel `bun:"table:subject"`
	ID            int64     `bun:"id,pk,autoincrement"`
	Name          string    `bun:"name"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
}
