package httpserver

import (
	"fmt"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
)

func toResponseTender(tender domain.Tender) TenderResponse {
	id := fmt.Sprintf("%d", tender.ID())
	return TenderResponse{
		ID:          id,
		Name:        tender.Name(),
		Description: tender.Description(),
		ServiceType: tender.ServiceType(),
		Status:      tender.Status(),
		Version:     tender.Version(),
		CreatedAt:   tender.CreatedAt(),
	}
}

func toResponseBid(bid domain.Bid) BidResponse {
	return BidResponse{
		ID:         bid.ID(),
		Name:       bid.Name(),
		Status:     bid.Status(),
		AuthorType: bid.AuthorType(),
		AuthorId:   bid.AuthorId(),
		Version:    bid.Version(),
		CreatedAt:  bid.CreatedAt(),
	}
}

func toDomainTender(tenderRequest TenderRequest) (domain.Tender, error) {
	return domain.NewTender(domain.NewTenderData{
		Name:           tenderRequest.Name,
		Description:    tenderRequest.Description,
		ServiceType:    tenderRequest.ServiceType,
		Status:         tenderRequest.Status,
		OrganizationId: tenderRequest.OrganizationId,
		UserId:         tenderRequest.UserID,
	})
}

func toDomainBid(bidRequest BidRequest) (domain.Bid, error) {
	return domain.NewBid(domain.NewBidData{
		Name:            bidRequest.Name,
		Description:     bidRequest.Description,
		TenderId:        bidRequest.TenderId,
		OrganizationId:  bidRequest.OrganizationId,
		CreatorUsername: bidRequest.CreatorUsername,
		AuthorId:        bidRequest.UserID,
	})
}

func toDomainUpdateTender(tenderRequest UpdateTenderRequest) (domain.Tender, error) {
	return domain.NewTender(domain.NewTenderData{
		Name:        tenderRequest.Name,
		Description: tenderRequest.Description,
		ServiceType: tenderRequest.ServiceType,
	})
}

func toDomainUpdateBid(bidRequest UpdateBidRequest) (domain.Bid, error) {
	return domain.NewBid(domain.NewBidData{
		Name:        bidRequest.Name,
		Description: bidRequest.Description,
	})
}
