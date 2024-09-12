package domain

import (
	"time"
)

type Tender struct {
	id             int
	name           string
	description    string
	serviceType    string
	status         string
	organizationId int
	version        int32
	createdAt      time.Time
	userId         int
}

type NewTenderData struct {
	ID             int
	Name           string
	Description    string
	ServiceType    string
	Status         string
	OrganizationId int
	Version        int32
	CreatedAt      time.Time
	UserId         int
}

// NewTender creates a new tender.
func NewTender(data NewTenderData) (Tender, error) {
	return Tender{
		id:             data.ID,
		name:           data.Name,
		description:    data.Description,
		serviceType:    data.ServiceType,
		status:         data.Status,
		organizationId: data.OrganizationId,
		version:        data.Version,
		createdAt:      data.CreatedAt,
		userId:         data.UserId,
	}, nil
}

func (t Tender) ID() int              { return t.id }
func (t Tender) Name() string         { return t.name }
func (t Tender) Description() string  { return t.description }
func (t Tender) ServiceType() string  { return t.serviceType }
func (t Tender) Status() string       { return t.status }
func (t Tender) OrganizationId() int  { return t.organizationId }
func (t Tender) Version() int32       { return t.version }
func (t Tender) CreatedAt() time.Time { return t.createdAt }
func (t Tender) UserID() int          { return t.userId }
