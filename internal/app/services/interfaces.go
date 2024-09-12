package services

import (
	"context"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, username string) (*domain.User, error)
	GetUserByID(ctx context.Context, id int) (domain.User, error)
}

type TenderRepository interface {
	GetTenders(ctx context.Context, serviceType []string, limit, offset int) ([]domain.Tender, error)
	GetTenderByID(ctx context.Context, tenderID int) (domain.Tender, error)
	CreateTender(ctx context.Context, tender domain.Tender) (domain.Tender, error)
	GetUserTenders(ctx context.Context, userID int, limit, offset int) ([]domain.Tender, error)
	GetTenderStatus(ctx context.Context, id int) (string, error)
	UpdateTenderStatus(ctx context.Context, id int, status string) (domain.Tender, error)
	UpdateTender(ctx context.Context, tender domain.Tender) (domain.Tender, error)
	RollbackVersion(ctx context.Context, tenderID, version int) (domain.Tender, error)
}

type BidRepository interface {
	CreateBid(ctx context.Context, b domain.Bid) (domain.Bid, error)
	GetBidByID(ctx context.Context, bidID int) (domain.Bid, error)
	GetUserBids(ctx context.Context, userID int, limit, offset int) ([]domain.Bid, error)
	GetBidStatus(ctx context.Context, id int) (string, error)
	GetTenderBids(ctx context.Context, tenderID, userID, limit, offset int) ([]domain.Bid, error)
	UpdateBid(ctx context.Context, bid domain.Bid) (domain.Bid, error)
	UpdateBidStatus(ctx context.Context, id int, status string) (domain.Bid, error)
	UpdateBidDescription(ctx context.Context, id int, desc string) (domain.Bid, error)
	RollbackBidVersion(ctx context.Context, id, version int) (domain.Bid, error)
	GetReviews(ctx context.Context, tenderID, userID, limit, offset int) ([]domain.Bid, error)
}
