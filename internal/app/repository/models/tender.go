package models

import (
	"time"

	"github.com/uptrace/bun"
)

// Tender is a domain tender.
type Tender struct {
	bun.BaseModel  `bun:"table:tender"`
	ID             int `bun:",pk,autoincrement"`
	Name           string
	Description    string
	ServiceType    string
	Status         string
	OrganizationId int
	UserId         int
	Version        int32     `bun:",default:1"`
	CreatedAt      time.Time `bun:",nullzero"`
}
