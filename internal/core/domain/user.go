package domain

type User struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}
