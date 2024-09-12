package pgrepo

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/repository/models"
)

func domainToUser(user domain.User) models.User {
	return models.User{
		ID:        user.ID(),
		Username:  user.Username(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
	}
}

func userToDomain(user models.User) (domain.User, error) {
	return domain.NewUser(domain.NewUserData{
		ID:             user.ID,
		Username:       user.Username,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		OrganizationID: user.OrganizationID,
	})
}

func domainToTender(tender domain.Tender) models.Tender {
	return models.Tender{
		ID:             tender.ID(),
		Name:           tender.Name(),
		Description:    tender.Description(),
		ServiceType:    tender.ServiceType(),
		Status:         tender.Status(),
		OrganizationId: tender.OrganizationId(),
		UserId:         tender.UserID(),
		Version:        tender.Version(),
		CreatedAt:      tender.CreatedAt(),
	}
}

func domainToBid(bid domain.Bid) models.Bid {
	return models.Bid{
		ID:          bid.ID(),
		Name:        bid.Name(),
		Description: bid.Description(),
		TenderId:    bid.TendedId(),
		Status:      bid.Status(),
		AuthorId:    bid.AuthorId(),
		AuthorType:  bid.AuthorType(),
		Version:     bid.Version(),
		CreatedAt:   bid.CreatedAt(),
	}
}

func bidToDomain(bid models.Bid) (domain.Bid, error) {
	return domain.NewBid(domain.NewBidData{
		ID:          bid.ID,
		Name:        bid.Name,
		Description: bid.Description,
		Status:      bid.Status,
		AuthorId:    bid.AuthorId,
		TenderId:    bid.TenderId,
		AuthorType:  bid.AuthorType,
		Version:     bid.Version,
	})
}

func tenderToDomain(tender models.Tender) (domain.Tender, error) {
	return domain.NewTender(domain.NewTenderData{
		ID:          tender.ID,
		Name:        tender.Name,
		Description: tender.Description,
		ServiceType: tender.ServiceType,
		Status:      tender.Status,
		Version:     tender.Version,
		CreatedAt:   tender.CreatedAt,
	})
}
