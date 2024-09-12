package domain

import (
	"time"
)

type Bid struct {
	id             int
	name           string
	description    string
	status         string
	tenderId       int
	organizationId int
	authorType     string
	authorId       int
	version        int32
	createdAt      time.Time
}

type NewBidData struct {
	ID              int
	Name            string
	Description     string
	Status          string
	TenderId        int
	OrganizationId  int
	CreatorUsername string
	AuthorType      string
	AuthorId        int
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

func (t Bid) ID() int              { return t.id }
func (t Bid) Name() string         { return t.name }
func (t Bid) Description() string  { return t.description }
func (t Bid) Status() string       { return t.status }
func (t Bid) AuthorType() string   { return t.authorType }
func (t Bid) AuthorId() int        { return t.authorId }
func (t Bid) TendedId() int        { return t.tenderId }
func (t Bid) OrganizationId() int  { return t.organizationId }
func (t Bid) Version() int32       { return t.version }
func (t Bid) CreatedAt() time.Time { return t.createdAt }

type BidResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	AuthorType string    `json:"authorType"`
	AuthorId   string    `json:"authorId"`
	Version    int32     `json:"version"`
	CreatedAt  time.Time `json:"createdAt"`
}
