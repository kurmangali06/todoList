package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"todolist/internal/core/domain"
	core_error "todolist/internal/core/errors"
	core_postgres_pool "todolist/internal/core/repository/postgres/pool"

	"github.com/google/uuid"
)

func (r *UsersRepository) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, full_name, phone_number
		FROM todoapp.users
		WHERE id = $1;
	`

	row := r.pool.QueryRow(ctx, query, id)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%s': %w",
				id,
				core_error.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.FullName,
		userModel.PhoneNumber,
		userModel.ID,
		userModel.Version,
	)

	return userDomain, nil
}
