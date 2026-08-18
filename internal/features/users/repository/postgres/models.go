package users_postgres_repository

import (
	"todolist/internal/core/domain"

	"github.com/google/uuid"
)

type UserModel struct {
	ID          uuid.UUID
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	usersDomain := make([]domain.User, len(users))

	for i, user := range users {
		usersDomain[i] = domain.NewUser(
			user.FullName,
			user.PhoneNumber,
			user.ID,
			user.Version,
		)
	}
	return usersDomain
}
