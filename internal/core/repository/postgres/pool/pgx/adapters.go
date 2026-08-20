package core_pgx_pool

import (
	"errors"
	"fmt"
	core_postgres_pool "todolist/internal/core/repository/postgres/pool"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRows struct {
	pgx.Rows
}

type pgxRow struct {
	pgx.Row
}

func (r pgxRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return core_postgres_pool.ErrNoRows
		}
		return err
	}

	return nil
}

type pgxCommandTag struct {
	pgconn.CommandTag
}

func mapErrors(err error) error {
	const (
		// Код PostgreSQL для ошибки нарушения внешнего ключа.
		// Полный список кодов: https://www.postgresql.org/docs/current/errcodes-appendix.html
		pgxViolatesForeignKeyErrorCode = "23503"
	)

	// pgx.ErrNoRows → наш ErrNoRows (запись не найдена)
	if errors.Is(err, pgx.ErrNoRows) {
		return core_postgres_pool.ErrNoRows
	}

	// Проверяем, является ли ошибка структурированной PostgreSQL-ошибкой.
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgxViolatesForeignKeyErrorCode {
			return fmt.Errorf(
				"%v: %w",
				err,
				core_postgres_pool.ErrViolatesForeignKey,
			)
		}
	}

	// Все остальные ошибки оборачиваем в ErrUnknown.
	return fmt.Errorf(
		"%v: %w",
		err,
		core_postgres_pool.ErrUnknown,
	)
}
