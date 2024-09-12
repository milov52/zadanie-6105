package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel  `bun:"table:employee"`
	ID             int `bun:",pk,autoincrement"`
	Username       string
	FirstName      string
	LastName       string
	OrganizationID int
	CreatedAt      time.Time `bun:",nullzero"`
	UpdatedAt      time.Time `bun:",nullzero"`
}
