package models

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:employee"`
	ID            int `bun:",pk,autoincrement"`
	Username      string
	FirstName     string
	LastName      string
	CreatedAt     time.Time `bun:",nullzero"`
	UpdatedAt     time.Time `bun:",nullzero"`
}
