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
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
}
