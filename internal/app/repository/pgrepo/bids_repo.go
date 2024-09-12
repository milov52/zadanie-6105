package pgrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/uptrace/bun"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/repository/models"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/pkg/pg"
)

type BidRepo struct {
	db *pg.DB
}

func NewBidRepo(db *pg.DB) *BidRepo {
	return &BidRepo{
		db: db,
	}
}

// Вспомогательная функция для проверки и добавления условий
func addWhereCondition(query *bun.SelectQuery, field string, params map[string]int) {
	if value, ok := params[field]; ok {
		query.Where(fmt.Sprintf("%s = ?", field), value)
	}
}

func (r BidRepo) CreateBid(ctx context.Context, b domain.Bid) (domain.Bid, error) {
	dbBid := domainToBid(b)

	var insertedBid models.Bid
	err := r.db.NewInsert().Model(&dbBid).Returning("*").Scan(ctx, &insertedBid)
	if err != nil {
		return domain.Bid{}, fmt.Errorf("failed to insert tender: %w", err)
	}

	domainBid, err := bidToDomain(insertedBid)
	if err != nil {
		return domain.Bid{}, fmt.Errorf("failed to create domain tender: %w", err)
	}

	return domainBid, nil
}

func (r BidRepo) getBids(ctx context.Context, params map[string]int) ([]domain.Bid, error) {
	var bids []models.Bid

	query := r.db.NewSelect().Model(&bids)

	addWhereCondition(query, "author_id", params)
	addWhereCondition(query, "tender_id", params)
	addWhereCondition(query, "id", params)

	if limit, ok := params["limit"]; ok && limit > 0 {
		query.Limit(limit)
	}

	if offset, ok := params["offset"]; ok && offset > 0 {
		query.Offset(offset)
	}

	query.Order("name")
	err := query.Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get tenders: %w", err)
	}

	domainBids := make([]domain.Bid, len(bids))
	for i, bid := range bids {
		domainBid, err := bidToDomain(bid)
		if err != nil {
			return nil, fmt.Errorf("failed to create domain tender: %w", err)
		}
		domainBids[i] = domainBid
	}

	return domainBids, nil
}

func (r BidRepo) GetUserBids(ctx context.Context, userID, limit, offset int) ([]domain.Bid, error) {
	params := map[string]int{
		"author_id": userID,
		"limit":     limit,
		"offset":    offset,
	}
	return r.getBids(ctx, params)
}

func (r BidRepo) GetTenderBids(ctx context.Context, tenderID, userID, limit, offset int) ([]domain.Bid, error) {
	params := map[string]int{
		"author_id": userID,
		"tender_id": tenderID,
		"limit":     limit,
		"offset":    offset,
	}
	return r.getBids(ctx, params)
}

func (r BidRepo) GetBidByID(ctx context.Context, bidID int) (domain.Bid, error) {
	params := map[string]int{
		"id": bidID,
	}
	bids, err := r.getBids(ctx, params)
	return bids[0], err
}

func (r BidRepo) GetBidStatus(ctx context.Context, id int) (string, error) {
	var status string
	err := r.db.NewSelect().
		Model((*models.Bid)(nil)).
		Column("status").
		Where("id = ?", id).
		Scan(ctx, &status)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", domain.ErrNotFound
		}
		return "", fmt.Errorf("failed to get bid status: %w", err)
	}

	return status, nil
}

func (r BidRepo) UpdateBidStatus(ctx context.Context, id int, status string) (domain.Bid, error) {
	var updatedBid models.Bid

	err := r.db.NewUpdate().
		Model((*models.Bid)(nil)).
		Set("status = ?", status).
		Where("id = ?", id).
		Returning("*").
		Scan(ctx, &updatedBid)
	if err != nil {
		return domain.Bid{}, fmt.Errorf("failed to update a status: %w", err)
	}

	domainBid, err := bidToDomain(updatedBid)
	if err != nil {
		return domain.Bid{}, fmt.Errorf("failed to create domain bid: %w", err)
	}

	return domainBid, nil
}

func (r BidRepo) UpdateBidDescription(ctx context.Context, id int, desc string) (domain.Bid, error) {
	var updatedBid models.Bid

	err := r.db.NewUpdate().
		Model((*models.Bid)(nil)).
		Set("description = ?", desc).
		Where("id = ?", id).
		Returning("*").
		Scan(ctx, &updatedBid)
	if err != nil {
		return domain.Bid{}, fmt.Errorf("failed to update a description: %w", err)
	}

	domainBid, err := bidToDomain(updatedBid)
	if err != nil {
		return domain.Bid{}, fmt.Errorf("failed to create domain bid: %w", err)
	}

	return domainBid, nil
}

func (r BidRepo) UpdateBid(ctx context.Context, bid domain.Bid) (domain.Bid, error) {
	dbBid := domainToBid(bid)

	var updatedBid models.Bid
	err := r.db.NewUpdate().
		Model(&dbBid).
		Where("id = ?", dbBid.ID).
		Set("version = version + 1").
		ExcludeColumn("id", "status", "tender_id", "author_id", "author_type", "created_at").
		Returning("*").
		Scan(ctx, &updatedBid)

	if err != nil {
		return domain.Bid{}, fmt.Errorf("failed to update a tender: %w", err)
	}

	domainBid, err := bidToDomain(updatedBid)
	if err != nil {
		return domain.Bid{}, fmt.Errorf("failed to create domain bid: %w", err)
	}

	return domainBid, nil
}

func (r BidRepo) RollbackBidVersion(ctx context.Context, bidID, version int) (domain.Bid, error) {
	var updatedBid models.Bid

	err := r.db.NewUpdate().
		Model((*models.Tender)(nil)).
		Set("version = ?", version).
		Where("id = ?", bidID).
		Returning("*").
		Scan(ctx, &updatedBid)
	if err != nil {
		return domain.Bid{}, fmt.Errorf("failed to update a version: %w", err)
	}

	domainBid, err := bidToDomain(updatedBid)
	if err != nil {
		return domain.Bid{}, fmt.Errorf("failed to create domain bid: %w", err)
	}

	return domainBid, nil
}

func (r BidRepo) GetReviews(ctx context.Context, tenderID, userID, limit, offset int) ([]domain.Bid, error) {
	params := map[string]int{
		"tender_id": tenderID,
		"author_id": userID,
		"limit":     limit,
		"offset":    offset,
	}
	return r.getBids(ctx, params)
}
