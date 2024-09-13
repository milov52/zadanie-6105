package domain

import "github.com/google/uuid"

// User is a domain User.
type User struct {
	id             uuid.UUID
	username       string
	firstName      string
	lastName       string
	organizationID uuid.UUID
}

type NewUserData struct {
	ID             uuid.UUID
	Username       string
	FirstName      string
	LastName       string
	OrganizationID uuid.UUID
}

// NewUser creates a new user.
func NewUser(data NewUserData) (User, error) {
	return User{
		id:             data.ID,
		username:       data.Username,
		firstName:      data.FirstName,
		lastName:       data.LastName,
		organizationID: data.OrganizationID,
	}, nil
}

func (u User) ID() uuid.UUID { return u.id }

func (u User) Username() string { return u.username }

func (u User) FirstName() string { return u.firstName }

func (u User) LastName() string { return u.lastName }

func (u User) OrganizationID() uuid.UUID { return u.organizationID }
