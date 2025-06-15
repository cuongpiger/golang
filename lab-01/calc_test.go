package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestAddWithTestify(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	// Using testify's assert package for better readability
	assert.Equal(t, expected, result, "Add(2, 3) should equal 5")
}

func TestBasicAssertions(t *testing.T) {
	// Equality assertions
	assert.Equal(t, 123, 123, "numbers should be equal")
	assert.NotEqual(t, 123, 456, "numbers should not be equal")

	// Boolean assertions
	assert.True(t, 1 < 2, "1 should be less than 2")
	assert.False(t, 1 > 2, "1 should not be greater than 2")

	// Nil assertions
	var nilPointer *int
	assert.Nil(t, nilPointer, "pointer should be nil")

	nonNilPointer := new(int)
	assert.NotNil(t, nonNilPointer, "pointer should not be nil")
}

// Run 'go test' in the terminal to execute this test
