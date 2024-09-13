package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Tender is a domain tender.
type Tender struct {
	bun.BaseModel  `bun:"table:tender"`
	ID             uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Name           string
	Description    string
	ServiceType    string
	Status         string    `bun:",default:Created"`
	OrganizationId uuid.UUID `bun:",type:uuid"`
	UserId         uuid.UUID `bun:",type:uuid"`
	Version        int32     `bun:",default:1"`
	CreatedAt      time.Time `bun:",nullzero"`
}
