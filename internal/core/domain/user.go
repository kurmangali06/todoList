package domain

import (
	"fmt"
	"regexp"
	core_error "todolist/internal/core/errors"

	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID
	Version int

	FullName    string
	PhoneNumber *string
}

func NewUser(
	fullName string,
	phoneNumber *string,
	ID uuid.UUID,
	Version int,
) User {
	return User{
		ID:          ID,
		Version:     Version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func (u User) Validate() error {
	fullNameLength := len([]rune(u.FullName))

	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalid `FullName` len: %d: %w",
			fullNameLength,
			core_error.ErrInvalidArgument,
		)
	}

	if u.PhoneNumber != nil {
		phoneNumberLength := len([]rune(*u.PhoneNumber))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf(
				"invalid `PhoneNumber` len: %d: %w",
				phoneNumberLength,
				core_error.ErrInvalidArgument,
			)
		}
		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid `PhoneNumber` format: %w",
				core_error.ErrInvalidArgument,
			)
		}
	}

	return nil
}
