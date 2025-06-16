package project

type User struct {
	ID   int
	Name string
}

type UserClient interface {
	GetUserByID(id int) (*User, error)
}
