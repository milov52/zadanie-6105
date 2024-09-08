package services

import (
	"context"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725732025-team-78758/zadanie-6105OD/internal/app/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, username string) (domain.User, error)
	GetUserByID(ctx context.Context, id int) (domain.User, error)
}
