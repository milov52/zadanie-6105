package httpserver

import (
	"context"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
)

// UserService is a user service
type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, username string) (domain.User, error)
	GetUserByID(ctx context.Context, id string) (domain.User, error)
}

type TenderService interface {
	GetTenders(ctx context.Context, serviceType []string, limit, offset int) ([]domain.Tender, error)
	GetUserTenders(ctx context.Context, userID string, limit, offset int) ([]domain.Tender, error)
	GetTenderByID(ctx context.Context, tenderID string) (domain.Tender, error)
	GetTenderStatus(ctx context.Context, id string) (string, error)
	UpdateTenderStatus(ctx context.Context, id string, status string) (domain.Tender, error)
	CreateTender(ctx context.Context, tender domain.Tender) (domain.Tender, error)
	UpdateTender(ctx context.Context, tender domain.Tender) (domain.Tender, error)
	RollbackVersion(ctx context.Context, tenderID string, version int) (domain.Tender, error)
}

type BidService interface {
	CreateBid(ctx context.Context, b domain.Bid) (domain.Bid, error)
	GetBidByID(ctx context.Context, bidID string) (domain.Bid, error)
	GetUserBids(ctx context.Context, userID string, limit, offset int) ([]domain.Bid, error)
	GetBidStatus(ctx context.Context, id string) (string, error)
	GetTenderBids(ctx context.Context, tenderID, userID string, limit, offset int) ([]domain.Bid, error)
	UpdateBid(ctx context.Context, bid domain.Bid) (domain.Bid, error)
	UpdateBidStatus(ctx context.Context, id, status string) (domain.Bid, error)
	UpdateBidDescription(ctx context.Context, id, desc string) (domain.Bid, error)
	RollbackBidVersion(ctx context.Context, id string, version int) (domain.Bid, error)
	GetReviews(ctx context.Context, tenderID, userID string, limit, offset int) ([]domain.Bid, error)
}
