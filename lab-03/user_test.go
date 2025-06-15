package lab03

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserCreation(t *testing.T) {
	u, err := CreateUser("Cuong. Duong Manh", "cuong.duong@email.com")
	
	// If user creation fails, no point continuing
    require.NoError(t, err, "user creation should not fail")
    require.NotNil(t, u, "user should not be nil")

    // These will only run if the above assertions pass
    require.Equal(t, "Cuong. Duong Manh", u.Name, "user name should match")
    require.Equal(t, "cuong.duong@email.com", u.Email, "user email should match")
}