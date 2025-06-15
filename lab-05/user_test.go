package user

import (
    "errors"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// Create a mock implementation of DataStore
type MockDataStore struct {
    mock.Mock
}

func (m *MockDataStore) GetUser(id string) (*User, error) {
    args := m.Called(id)

    // If the first return value is nil, return nil, args.Error(1)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }

    // Otherwise return the user
    return args.Get(0).(*User), args.Error(1)
}

func (m *MockDataStore) SaveUser(user *User) error {
    args := m.Called(user)
    return args.Error(0)
}

func TestUserService_GetUserByID(t *testing.T) {
    // Create a mock data store
    mockStore := new(MockDataStore)

    // Create a user service with the mock store
    service := NewUserService(mockStore)

    // Setup expectations
    testUser := &User{ID: "123", Name: "John", Email: "john@example.com"}
    mockStore.On("GetUser", "123").Return(testUser, nil)

    // Test the service
    user, err := service.GetUserByID("123")

    // Assertions
    assert.NoError(t, err, "should not return error")
    assert.Equal(t, testUser, user, "should return the user")

    // Verify expectations were met
    mockStore.AssertExpectations(t)
}

func TestUserService_UpdateUserEmail(t *testing.T) {
    // Create a mock data store
    mockStore := new(MockDataStore)

    // Create a user service with the mock store
    service := NewUserService(mockStore)

    // Setup expectations
    testUser := &User{ID: "123", Name: "John", Email: "john@example.com"}

    mockStore.On("GetUser", "123").Return(testUser, nil)
    mockStore.On("SaveUser", mock.MatchedBy(func(u *User) bool {
        return u.ID == "123" && u.Email == "newemail@example.com"
    })).Return(nil)

    // Test the service
    err := service.UpdateUserEmail("123", "newemail@example.com")

    // Assertions
    assert.NoError(t, err, "should not return error")

    // Verify expectations were met
    mockStore.AssertExpectations(t)
}

func TestUserService_UpdateUserEmail_Error(t *testing.T) {
    // Create a mock data store
    mockStore := new(MockDataStore)

    // Create a user service with the mock store
    service := NewUserService(mockStore)

    // Setup expectations - user not found
    mockStore.On("GetUser", "999").Return(nil, errors.New("user not found"))

    // Test the service
    err := service.UpdateUserEmail("999", "newemail@example.com")

    // Assertions
    assert.Error(t, err, "should return error")
    assert.Equal(t, "user not found", err.Error(), "error message should match")

    // Verify expectations were met
    mockStore.AssertExpectations(t)
}