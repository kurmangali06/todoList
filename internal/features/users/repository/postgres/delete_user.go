package users_postgres_repository

import (
	"context"
	"fmt"
	core_error "todolist/internal/core/errors"

	"github.com/google/uuid"
)

func (r *UsersRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		DELETE FROM todoapp.users
		WHERE id = $1;
	`
	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("Exec query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id='%s': %w", id, core_error.ErrNotFound)
	}

	return nil
}
