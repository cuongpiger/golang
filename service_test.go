package project

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/cuongpiger/golang/mocks"
)

func TestUserService_GetUserDetails_Success(t *testing.T) {
	// Create a new mock client
	mockClient := new(mocks.UserClient)

	// Define what the mock should return when `GetUserByID` is called
	mockClient.On("GetUserByID", 1).Return(&User{
		ID:   1,
		Name: "John Doe",
	}, nil)

	// Create the UserService with the mock client
	service := NewUserService(mockClient)

	// Test the GetUserDetails method
	result, err := service.GetUserDetails(1)

	// Use `require` for error checks
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// Use `assert` for value checks
	assert.Equal(t, "User: John Doe (ID: 1)", result)

	// Ensure that the `GetUserByID` method was called exactly once
	mockClient.AssertExpectations(t)
}

func TestUserService_GetUserDetails_Error(t *testing.T) {
	// Create a new mock client
	mockClient := new(mocks.UserClient)

	// Define what the mock should return when `GetUserByID` is called with an error
	mockClient.On("GetUserByID", 2).Return(nil, errors.New("user not found"))

	// Create the UserService with the mock client
	service := NewUserService(mockClient)

	// Test the GetUserDetails method
	result, err := service.GetUserDetails(2)

	// Use `require` for error checks
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")

	// Ensure that the result is empty
	assert.Empty(t, result)

	// Ensure that the `GetUserByID` method was called exactly once
	mockClient.AssertExpectations(t)
}
