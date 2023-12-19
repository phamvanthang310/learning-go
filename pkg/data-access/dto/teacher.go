package dto

import (
	"github.com/uptrace/bun"
	"time"
)

type Teacher struct {
	bun.BaseModel `bun:"table:teacher"`
	ID            int64     `bun:"pk,autoincrement"`
	Name          string    `bun:"name"`
	Username      string    `bun:"username"`
	Password      string    `bun:"password"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
}
