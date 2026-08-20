package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"todolist/internal/core/domain"
	core_error "todolist/internal/core/errors"
	core_postgres_pool "todolist/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE todoapp.users
		SET 
		    full_name = $1, 
		    phone_number = $2,
			version = version + 1
		WHERE id = $3 AND version = $4
		RETURNING id, version, full_name, phone_number;
	`
	row := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber, user.ID, user.Version)
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
				"user with id='%s' concurrently accessed: %w",
				user.ID,
				core_error.ErrConflict,
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
