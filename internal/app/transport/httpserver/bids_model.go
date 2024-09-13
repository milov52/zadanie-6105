package httpserver

import (
	"fmt"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
	"time"
)

type BidRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Status          string `json:"status"`
	TenderId        string `json:"tenderId"`
	OrganizationId  string `json:"organizationId"`
	CreatorUsername string `json:"creatorUsername"`
	UserID          string
}

type BidResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	AuthorType string    `json:"authorType"`
	AuthorId   string    `json:"authorId"`
	Version    int32     `json:"version"`
	CreatedAt  time.Time `json:"createdAt"`
}

type UpdateBidRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

const (
	BidStatusTypeCreated   = "Created"
	BidStatusTypePublished = "Published"
	BidStatusTypeCanceled  = "Canceled"
	BidStatusTypeApproved  = "Approved"
	BidStatusTypeRejected  = "Rejected"
)

func validateBidStatus(status string) error {
	switch status {
	case BidStatusTypeCreated, BidStatusTypePublished, BidStatusTypeCanceled, BidStatusTypeApproved, BidStatusTypeRejected:
		return nil
	default:
		return fmt.Errorf("%w: tender status is invalid", domain.ErrNegative)
	}
}

func validateDecision(decision string) error {
	switch decision {
	case BidStatusTypeApproved, BidStatusTypeRejected:
		return nil
	default:
		return fmt.Errorf("%w: decision status is invalid", domain.ErrNegative)
	}
}

func (r *BidRequest) Validate() error {
	if r.Name == "" || len(r.Name) > 100 {
		return fmt.Errorf("%w: name", domain.ErrNegative)
	}
	if r.Description == "" || len(r.Name) > 500 {
		return fmt.Errorf("%w: description", domain.ErrNegative)
	}
	if r.TenderId == "" || len(r.Name) > 100 {
		return fmt.Errorf("%w: organizationId", domain.ErrNegative)
	}

	if err := validateBidStatus(r.Status); err != nil {
		return err
	}
	return nil
}

func (ru *UpdateBidRequest) Validate() error {
	if ru.Name == "" || len(ru.Name) > 100 {
		return fmt.Errorf("%w: name", domain.ErrNegative)
	}
	if ru.Description == "" || len(ru.Name) > 500 {
		return fmt.Errorf("%w: description", domain.ErrNegative)
	}
	return nil
}
