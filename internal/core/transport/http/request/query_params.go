package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"
	core_error "todolist/internal/core/errors"

	"github.com/google/uuid"
)

func GetUUIDQueryParam(r *http.Request, key string) (*uuid.UUID, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := uuid.Parse(param)
	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key='%s' not a valid uuid: %v: %w",
			param,
			key,
			err,
			core_error.ErrInvalidArgument,
		)
	}

	return &val, nil
}

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key='%s' not a valid integer: %v: %w",
			param,
			key,
			err,
			core_error.ErrInvalidArgument,
		)
	}

	return &val, nil
}
