package httpserver

import (
	"fmt"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
	"github.com/google/uuid"
	"time"
)

type TenderRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ServiceType     string `json:"serviceType"`
	Status          string `json:"status"`
	OrganizationId  string `json:"organizationId"`
	CreatorUsername string `json:"creatorUsername"`
	UserID          string
}

type TenderResponse struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ServiceType    string    `json:"serviceType"`
	Status         string    `json:"status"`
	OrganizationId string    `json:"organizationId"`
	Version        int32     `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
}

type UpdateTenderRequest struct {
	ID          uuid.UUID
	Name        string `json:"name"`
	Description string `json:"description"`
	ServiceType string `json:"serviceType"`
}

const (
	ServiceTypeConstruction = "Construction"
	ServiceTypeDelivery     = "Delivery"
	ServiceTypeManufacture  = "Manufacture"
)

const (
	StatusTypeCreated   = "Created"
	StatusTypePublished = "Published"
	StatusTypeClosed    = "Closed"
)

func validateTenderServiceType(serviceType string) error {
	switch serviceType {
	case ServiceTypeConstruction, ServiceTypeDelivery, ServiceTypeManufacture:
		return nil
	default:
		return fmt.Errorf("%w: tender ServiceType is invalid", domain.ErrNegative)
	}
}

func validateTenderStatus(status string) error {
	switch status {
	case StatusTypeCreated, StatusTypePublished, StatusTypeClosed:
		return nil
	default:
		return fmt.Errorf("%w: tender status is invalid", domain.ErrNegative)
	}
}

func (r *TenderRequest) Validate() error {
	if r.Name == "" || len(r.Name) > 100 {
		return fmt.Errorf("%w: name", domain.ErrNegative)
	}
	if r.Description == "" || len(r.Name) > 500 {
		return fmt.Errorf("%w: description", domain.ErrNegative)
	}
	if err := validateTenderServiceType(r.ServiceType); err != nil {
		return err
	}

	if r.OrganizationId == "" || len(r.Name) > 100 {
		return fmt.Errorf("%w: organizationId", domain.ErrNegative)
	}
	return nil
}

func (ru *UpdateTenderRequest) Validate() error {
	if ru.Name == "" || len(ru.Name) > 100 {
		return fmt.Errorf("%w: name", domain.ErrNegative)
	}
	if ru.Description == "" || len(ru.Name) > 500 {
		return fmt.Errorf("%w: description", domain.ErrNegative)
	}
	if err := validateTenderServiceType(ru.ServiceType); err != nil {
		return err
	}
	return nil
}
