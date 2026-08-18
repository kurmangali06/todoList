package users_postgres_repository

import "github.com/google/uuid"

type UserModel struct {
	ID          uuid.UUID
	Version     int
	FullName    string
	PhoneNumber *string
}
