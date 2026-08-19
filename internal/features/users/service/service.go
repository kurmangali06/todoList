package users_service

import (
	"context"
	"todolist/internal/core/domain"

	"github.com/google/uuid"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		ID uuid.UUID,
	) (domain.User, error)
	DeleteUser(
		ctx context.Context,
		ID uuid.UUID,
	) error
	PatchUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersService(usersRepository UsersRepository) *UsersService {
	return &UsersService{usersRepository: usersRepository}
}
