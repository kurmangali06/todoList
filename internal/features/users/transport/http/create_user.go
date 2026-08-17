package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"
	core_logger "todolist/internal/core/logger"
)

type CreateUserRequest struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type CreateUserResponse struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Version     int    `json:"version"`
}

func (h *UserHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	log.Debug("create user")

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("error:", err)
	}
	rw.WriteHeader(http.StatusOK)
}
