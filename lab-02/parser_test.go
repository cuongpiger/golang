package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	// Test success case
	result, err := Parse("valid input")
	assert.NoError(t, err, "Expected no error for valid input")
	assert.Equal(t, 42, result, "Expected result to be 42 for valid input")

	// Test error case
	result, err = Parse("")
	assert.Error(t, err, "empty input should cause an error")
	assert.Equal(t, "empty input", err.Error(), "error message should match")
	assert.Equal(t, 0, result, "error case should return zero value")
}
