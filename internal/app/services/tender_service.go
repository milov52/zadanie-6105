package services

import (
	"context"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
)

type TenderService struct {
	repo TenderRepository
}

func NewTenderService(repo TenderRepository) TenderService {
	return TenderService{repo: repo}
}

func (s TenderService) CreateTender(ctx context.Context, tender domain.Tender) (domain.Tender, error) {
	return s.repo.CreateTender(ctx, tender)
}

func (s TenderService) GetTenders(ctx context.Context, serviceType []string, limit, offset int) ([]domain.Tender, error) {
	return s.repo.GetTenders(ctx, serviceType, limit, offset)
}

func (s TenderService) GetUserTenders(ctx context.Context, userID string, limit, offset int) ([]domain.Tender, error) {
	return s.repo.GetUserTenders(ctx, userID, limit, offset)
}

func (s TenderService) GetTenderStatus(ctx context.Context, id string) (string, error) {
	return s.repo.GetTenderStatus(ctx, id)
}

func (s TenderService) GetTenderByID(ctx context.Context, tenderID string) (domain.Tender, error) {
	return s.repo.GetTenderByID(ctx, tenderID)
}

func (s TenderService) UpdateTender(ctx context.Context, tender domain.Tender) (domain.Tender, error) {
	return s.repo.UpdateTender(ctx, tender)
}

func (s TenderService) UpdateTenderStatus(ctx context.Context, id, status string) (domain.Tender, error) {
	return s.repo.UpdateTenderStatus(ctx, id, status)
}

func (s TenderService) RollbackVersion(ctx context.Context, tenderID string, version int) (domain.Tender, error) {
	return s.repo.RollbackVersion(ctx, tenderID, version)
}
