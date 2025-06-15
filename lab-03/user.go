package lab03

type User struct {
	Name  string
	Email string
}

func CreateUser(name, email string) (User, error) {
	return User{Name: name, Email: email}, nil
}
