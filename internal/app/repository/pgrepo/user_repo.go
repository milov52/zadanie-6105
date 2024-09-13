package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/repository/models"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/pkg/pg"
)

type UserRepo struct {
	db *pg.DB
}

func NewUserRepo(db *pg.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r UserRepo) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	dbUser := domainToUser(user)

	var insertedUser models.User
	err := r.db.NewInsert().Model(&dbUser).Returning("*").Scan(ctx, &insertedUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to insert user: %w", err)
	}

	domainUser, err := userToDomain(insertedUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create domain user: %w", err)
	}

	return domainUser, nil
}

func (r UserRepo) GetUser(ctx context.Context, username string) (*domain.User, error) {
	var dbUser models.User

	err := r.db.NewSelect().
		Model((*models.User)(nil)).
		Column("id", "organization_responsible.organization_id").
		Join(`LEFT JOIN organization_responsible ON organization_responsible.user_id = "user".id`).
		Where(`"user".username = ?`, username).
		Scan(ctx, &dbUser)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	user, err := userToDomain(dbUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create domain user: %w", err)
	}

	return &user, nil
}

func (r UserRepo) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	var dbUser models.User
	err := r.db.NewSelect().
		Model((*models.User)(nil)).
		Column("id", "organization_responsible.organization_id").
		Join(`LEFT JOIN organization_responsible ON organization_responsible.user_id = "user".id`).
		Where(`"user".id = ?`, id).
		Scan(ctx, &dbUser)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrNotFound
		}
		return domain.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	user, err := userToDomain(dbUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create domain user: %w", err)
	}
	return user, nil
}
