package project

import "fmt"

type UserService struct {
	client UserClient
}

func NewUserService(client UserClient) *UserService {
	return &UserService{client: client}
}

func (s *UserService) GetUserDetails(id int) (string, error) {
	user, err := s.client.GetUserByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	return fmt.Sprintf("User: %s (ID: %d)", user.Name, user.ID), nil
}
