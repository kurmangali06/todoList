package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"
	core_error "todolist/internal/core/errors"

	"github.com/google/uuid"
)

func GetUUIDPathValue(r *http.Request, key string) (uuid.UUID, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return uuid.UUID{}, fmt.Errorf(
			"no key='%s' in path values: %w",
			key,
			core_error.ErrInvalidArgument,
		)
	}

	val, err := uuid.Parse(pathValue)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf(
			"path value='%s' by key='%s' not a valid uuid: %v: %w",
			pathValue,
			key,
			err,
			core_error.ErrInvalidArgument,
		)
	}

	return val, nil
}
func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)

	if pathValue == "" {
		return 0, fmt.Errorf(
			"no key='%s' in path values: %w",
			key,
			core_error.ErrInvalidArgument,
		)
	}

	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf(
			"path value='%s' by key='%s' not a valid integer: %v  : %w",
			pathValue,
			key,
			err,
			core_error.ErrInvalidArgument,
		)
	}

	return val, nil

}
