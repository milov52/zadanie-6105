package domain

// User is a domain User.
type User struct {
	id             int
	username       string
	firstName      string
	lastName       string
	organizationID int
}

type NewUserData struct {
	ID             int
	Username       string
	FirstName      string
	LastName       string
	OrganizationID int
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

func (u User) ID() int { return u.id }

func (u User) Username() string { return u.username }

func (u User) FirstName() string { return u.firstName }

func (u User) LastName() string { return u.lastName }

func (u User) OrganizationID() int { return u.organizationID }
