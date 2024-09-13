package domain

import (
	"time"

	"github.com/google/uuid"
)

type Bid struct {
	id             uuid.UUID
	name           string
	description    string
	status         string
	tenderId       uuid.UUID
	organizationId uuid.UUID
	authorType     string
	authorId       uuid.UUID
	version        int32
	createdAt      time.Time
}

type NewBidData struct {
	ID              uuid.UUID
	Name            string
	Description     string
	Status          string
	TenderId        uuid.UUID
	OrganizationId  uuid.UUID
	CreatorUsername string
	AuthorType      string
	AuthorId        uuid.UUID
	Version         int32
	CreatedAt       time.Time
}

// NewTender creates a new tender.
func NewBid(data NewBidData) (Bid, error) {
	return Bid{
		id:         data.ID,
		name:       data.Name,
		status:     data.Status,
		authorType: data.AuthorType,
		authorId:   data.AuthorId,
		version:    data.Version,
		createdAt:  data.CreatedAt,
	}, nil
}

func (t Bid) ID() uuid.UUID             { return t.id }
func (t Bid) Name() string              { return t.name }
func (t Bid) Description() string       { return t.description }
func (t Bid) Status() string            { return t.status }
func (t Bid) AuthorType() string        { return t.authorType }
func (t Bid) AuthorId() uuid.UUID       { return t.authorId }
func (t Bid) TendedId() uuid.UUID       { return t.tenderId }
func (t Bid) OrganizationId() uuid.UUID { return t.organizationId }
func (t Bid) Version() int32            { return t.version }
func (t Bid) CreatedAt() time.Time      { return t.createdAt }

type BidResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	AuthorType string    `json:"authorType"`
	AuthorId   string    `json:"authorId"`
	Version    int32     `json:"version"`
	CreatedAt  time.Time `json:"createdAt"`
}
