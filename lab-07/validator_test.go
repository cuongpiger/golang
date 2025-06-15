package validator

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func IsValidEmail(email string) bool {
    // Simple validation for demonstration
    return len(email) > 5 && contains(email, "@") && contains(email, ".")
}

func contains(s string, substr string) bool {
    for i := 0; i < len(s); i++ {
        if i+len(substr) <= len(s) && s[i:i+len(substr)] == substr {
            return true
        }
    }
    return false
}

func TestIsValidEmail(t *testing.T) {
    tests := []struct {
        name     string
        email    string
        expected bool
    }{
        {"valid email", "user@example.com", true},
        {"missing @", "userexample.com", false},
        {"missing .", "user@examplecom", false},
        {"too short", "u@e.c", false},
        {"empty string", "", false},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := IsValidEmail(test.email)
            assert.Equal(t, test.expected, result,
                        "IsValidEmail(%q) should be %v", test.email, test.expected)
        })
    }
}