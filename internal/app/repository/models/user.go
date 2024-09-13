package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel  `bun:"table:employee"`
	ID             uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Username       string
	FirstName      string
	LastName       string
	OrganizationID uuid.UUID `bun:",type:uuid"`
	CreatedAt      time.Time `bun:",nullzero"`
	UpdatedAt      time.Time `bun:",nullzero"`
}
