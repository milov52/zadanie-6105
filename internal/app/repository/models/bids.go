package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Tender is a domain tender.
type Bid struct {
	bun.BaseModel `bun:"table:bid"`
	ID            uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Name          string
	Description   string
	Status        string    `bun:",default:Created"`
	TenderId      uuid.UUID `bun:",type:uuid"`
	AuthorType    string
	AuthorId      uuid.UUID `bun:",type:uuid"`
	Version       int32     `bun:",default:1"`
	CreatedAt     time.Time `bun:",nullzero"`
}
