package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/uptrace/bun"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/repository/models"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/pkg/pg"
)

type TenderRepo struct {
	db *pg.DB
}

func NewTenderRepo(db *pg.DB) *TenderRepo {
	return &TenderRepo{
		db: db,
	}
}

func (r TenderRepo) CreateTender(ctx context.Context, tender domain.Tender) (domain.Tender, error) {
	dbTender := domainToTender(tender)

	var insertedTender models.Tender
	err := r.db.NewInsert().Model(&dbTender).Returning("*").Scan(ctx, &insertedTender)
	if err != nil {
		return domain.Tender{}, fmt.Errorf("failed to insert tender: %w", err)
	}

	domainTender, err := tenderToDomain(insertedTender)
	if err != nil {
		return domain.Tender{}, fmt.Errorf("failed to create domain tender: %w", err)
	}

	return domainTender, nil
}

func (r TenderRepo) GetTenders(ctx context.Context, serviceType []string, limit, offset int) ([]domain.Tender, error) {
	var tenders []models.Tender
	const statusClosed = "Closed"
	query := r.db.NewSelect().Model(&tenders)
	if len(serviceType) > 0 {
		query = query.Where("service_type IN (?) and status != ?", bun.In(serviceType), statusClosed)
	}
	if limit > 0 {
		query.Limit(limit)
	}
	if offset > 0 {
		query.Offset(offset)
	}
	query.Order("name")
	err := query.Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get tenders: %w", err)
	}

	domainTenders := make([]domain.Tender, len(tenders))
	for i, tender := range tenders {
		domainTender, err := tenderToDomain(tender)
		if err != nil {
			return nil, fmt.Errorf("failed to create domain tender: %w", err)
		}

		domainTenders[i] = domainTender
	}

	return domainTenders, nil
}

func (r TenderRepo) GetTenderByID(ctx context.Context, tenderID string) (domain.Tender, error) {
	var tender models.Tender
	const statusClosed = "Closed"
	query := r.db.NewSelect().Model(&tender).
		Where("id = ? and status != ?", tenderID, statusClosed)
	query.Order("name")
	err := query.Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Tender{}, domain.ErrNotFound
		}
		return domain.Tender{}, fmt.Errorf("failed to get tenders: %w", err)
	}

	domainTender, err := tenderToDomain(tender)
	return domainTender, nil
}

func (r TenderRepo) GetUserTenders(ctx context.Context, userID string, limit, offset int) ([]domain.Tender, error) {
	var tenders []models.Tender

	query := r.db.NewSelect().Model(&tenders)
	query = query.Where("user_id = (?)", userID)

	if limit > 0 {
		query.Limit(limit)
	}
	if offset > 0 {
		query.Offset(offset)
	}
	query.Order("name")
	err := query.Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get tenders: %w", err)
	}

	domainTenders := make([]domain.Tender, len(tenders))
	for i, tender := range tenders {
		domainTender, err := tenderToDomain(tender)
		if err != nil {
			return nil, fmt.Errorf("failed to create domain tender: %w", err)
		}
		domainTenders[i] = domainTender
	}

	return domainTenders, nil
}

func (r TenderRepo) GetTenderStatus(ctx context.Context, id string) (string, error) {
	var status string
	err := r.db.NewSelect().
		Model((*models.Tender)(nil)).
		Column("status").
		Where("id = ?", id).
		Scan(ctx, &status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", domain.ErrNotFound
		}
		return "", fmt.Errorf("failed to get tender status: %w", err)
	}

	return status, nil
}

func (r TenderRepo) UpdateTenderStatus(ctx context.Context, id, status string) (domain.Tender, error) {
	var updatedTender models.Tender

	err := r.db.NewUpdate().
		Model((*models.Tender)(nil)).
		Set("status = ?", status).
		Where("id = ?", id).
		Returning("*").
		Scan(ctx, &updatedTender)
	if err != nil {
		return domain.Tender{}, fmt.Errorf("failed to update a status: %w", err)
	}

	domainTender, err := tenderToDomain(updatedTender)
	if err != nil {
		return domain.Tender{}, fmt.Errorf("failed to create domain tender: %w", err)
	}

	return domainTender, nil
}

func (r TenderRepo) UpdateTender(ctx context.Context, tender domain.Tender) (domain.Tender, error) {
	dbTender := domainToTender(tender)

	var updatedTender models.Tender
	err := r.db.NewUpdate().
		Model(&dbTender).
		Where("id = ?", dbTender.ID).
		Set("version = version + 1").
		Set("name = ?", dbTender.Name).
		Set("description = ?", dbTender.Description).
		Set("service_type = ?", dbTender.ServiceType).
		//ExcludeColumn("status", "organization_id", "version", "created_at").
		Returning("*").
		Scan(ctx, &updatedTender)

	if err != nil {
		return domain.Tender{}, fmt.Errorf("failed to update a tender: %w", err)
	}

	domainTender, err := tenderToDomain(updatedTender)
	if err != nil {
		return domain.Tender{}, fmt.Errorf("failed to create domain book: %w", err)
	}

	return domainTender, nil
}

func (r TenderRepo) RollbackVersion(ctx context.Context, tenderID string, version int) (domain.Tender, error) {
	var updatedTender models.Tender

	err := r.db.NewUpdate().
		Model((*models.Tender)(nil)).
		Set("version = ?", version).
		Where("id = ?", tenderID).
		Returning("*").
		Scan(ctx, &updatedTender)
	if err != nil {
		return domain.Tender{}, fmt.Errorf("failed to update a version: %w", err)
	}

	domainTender, err := tenderToDomain(updatedTender)
	if err != nil {
		return domain.Tender{}, fmt.Errorf("failed to create domain tender: %w", err)
	}

	return domainTender, nil
}
