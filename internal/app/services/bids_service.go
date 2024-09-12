package services

import (
	"golang.org/x/net/context"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
)

type BidService struct {
	repo BidRepository
}

func NewBidService(repo BidRepository) BidService {
	return BidService{repo: repo}
}

func (s BidService) CreateBid(ctx context.Context, bid domain.Bid) (domain.Bid, error) {
	return s.repo.CreateBid(ctx, bid)
}

func (s BidService) GetUserBids(ctx context.Context, userID int, limit, offset int) ([]domain.Bid, error) {
	return s.repo.GetUserBids(ctx, userID, limit, offset)
}

func (s BidService) GetTenderBids(ctx context.Context, tenderID, userID, limit, offset int) ([]domain.Bid, error) {
	return s.repo.GetTenderBids(ctx, tenderID, userID, limit, offset)
}

func (s BidService) GetBidStatus(ctx context.Context, id int) (string, error) {
	return s.repo.GetBidStatus(ctx, id)
}

func (s BidService) GetBidByID(ctx context.Context, bidID int) (domain.Bid, error) {
	return s.repo.GetBidByID(ctx, bidID)
}

func (s BidService) UpdateBid(ctx context.Context, tender domain.Bid) (domain.Bid, error) {
	return s.repo.UpdateBid(ctx, tender)
}

func (s BidService) UpdateBidStatus(ctx context.Context, id int, status string) (domain.Bid, error) {
	return s.repo.UpdateBidStatus(ctx, id, status)
}

func (s BidService) UpdateBidDescription(ctx context.Context, id int, desc string) (domain.Bid, error) {
	return s.repo.UpdateBidDescription(ctx, id, desc)
}

func (s BidService) RollbackBidVersion(ctx context.Context, id, version int) (domain.Bid, error) {
	return s.repo.RollbackBidVersion(ctx, id, version)
}

func (s BidService) GetReviews(ctx context.Context, tenderID, userID, limit, offset int) ([]domain.Bid, error) {
	return s.repo.GetReviews(ctx, tenderID, userID, limit, offset)
}
