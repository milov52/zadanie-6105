package models

import (
	"time"

	"github.com/uptrace/bun"
)

// Tender is a domain tender.
type Bid struct {
	bun.BaseModel `bun:"table:bid"`
	ID            int `bun:",pk,autoincrement"`
	Name          string
	Description   string
	Status        string
	TenderId      int
	AuthorType    string
	AuthorId      int
	Version       int32     `bun:",default:1"`
	CreatedAt     time.Time `bun:",nullzero"`
}
