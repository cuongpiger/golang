package user

type DataStore interface {
    GetUser(id string) (*User, error)
    SaveUser(user *User) error
}

type User struct {
    ID    string
    Name  string
    Email string
}

type UserService struct {
    store DataStore
}

func NewUserService(store DataStore) *UserService {
    return &UserService{store: store}
}

func (s *UserService) GetUserByID(id string) (*User, error) {
    return s.store.GetUser(id)
}

func (s *UserService) UpdateUserEmail(id, newEmail string) error {
    user, err := s.store.GetUser(id)
    if err != nil {
        return err
    }

    user.Email = newEmail
    return s.store.SaveUser(user)
}