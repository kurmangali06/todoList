package users_postgres_repository

import (
	"context"
	"fmt"
	"todolist/internal/core/domain"
)

func (r *UsersRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO todoapp.users (full_name, phone_number)
		VALUES ($1, $2)
		RETURNING id, version, full_name, phone_number
	`
	row := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber)
	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)

	if err != nil {
		return domain.User{}, fmt.Errorf("scan user: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.FullName,
		userModel.PhoneNumber,
		userModel.ID,
		userModel.Version,
	)

	return userDomain, nil
}
