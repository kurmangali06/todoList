package users_service

import (
	"context"
	"fmt"
	"todolist/internal/core/domain"

	"github.com/google/uuid"
)

func (s *UsersService) PatchUser(
	ctx context.Context,
	ID uuid.UUID,
	patch domain.UserPatch,
) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, ID)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("apply patch to user: %w", err)
	}

	patchedUser, err := s.usersRepository.PatchUser(ctx, user)

	if err != nil {
		return domain.User{}, fmt.Errorf("update user in repository: %w", err)
	}

	return patchedUser, nil

}
