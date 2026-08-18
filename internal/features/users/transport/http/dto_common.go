package users_transport_http

import (
	"todolist/internal/core/domain"

	"github.com/google/uuid"
)

type UserDTOResponse struct {
	ID          uuid.UUID `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber *string   `json:"phone_number"`
	Version     int       `json:"version"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}
func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
