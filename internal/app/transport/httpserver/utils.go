package httpserver

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
	"github.com/google/uuid"
)

func toResponseTender(tender domain.Tender) TenderResponse {
	return TenderResponse{
		ID:             tender.ID().String(),
		Name:           tender.Name(),
		Description:    tender.Description(),
		ServiceType:    tender.ServiceType(),
		OrganizationId: tender.OrganizationId().String(),
		Status:         tender.Status(),
		Version:        tender.Version(),
		CreatedAt:      tender.CreatedAt(),
	}
}

func toResponseBid(bid domain.Bid) BidResponse {
	return BidResponse{
		ID:          bid.ID().String(),
		Name:        bid.Name(),
		Description: bid.Description(),
		Status:      bid.Status(),
		TenderId:    bid.TendedId().String(),
		AuthorType:  bid.AuthorType(),
		AuthorId:    bid.AuthorId().String(),
		Version:     bid.Version(),
		CreatedAt:   bid.CreatedAt(),
	}
}

func toDomainTender(tenderRequest TenderRequest) (domain.Tender, error) {
	organizationID, err := uuid.Parse(tenderRequest.OrganizationId)
	if err != nil {
		return domain.Tender{}, err
	}
	userId, err := uuid.Parse(tenderRequest.UserID)
	if err != nil {
		return domain.Tender{}, err
	}

	return domain.NewTender(domain.NewTenderData{
		Name:           tenderRequest.Name,
		Description:    tenderRequest.Description,
		ServiceType:    tenderRequest.ServiceType,
		Status:         tenderRequest.Status,
		OrganizationId: organizationID,
		UserId:         userId,
	})
}

func toDomainBid(bidRequest BidRequest) (domain.Bid, error) {
	tenderId, err := uuid.Parse(bidRequest.TenderId)
	if err != nil {
		return domain.Bid{}, err
	}
	userId, err := uuid.Parse(bidRequest.AuthorId)
	if err != nil {
		return domain.Bid{}, err
	}

	return domain.NewBid(domain.NewBidData{
		Name:        bidRequest.Name,
		Description: bidRequest.Description,
		TenderId:    tenderId,
		AuthorType:  bidRequest.AuthorType,
		AuthorId:    userId,
	})
}

func toDomainUpdateTender(tenderRequest UpdateTenderRequest) (domain.Tender, error) {
	return domain.NewTender(domain.NewTenderData{
		ID:          tenderRequest.ID,
		Name:        tenderRequest.Name,
		Description: tenderRequest.Description,
		ServiceType: tenderRequest.ServiceType,
	})
}

func toDomainUpdateBid(bidRequest UpdateBidRequest) (domain.Bid, error) {
	bidID, err := uuid.Parse(bidRequest.ID)
	if err != nil {
		return domain.Bid{}, err
	}
	return domain.NewBid(domain.NewBidData{
		ID:          bidID,
		Name:        bidRequest.Name,
		Description: bidRequest.Description,
	})
}
