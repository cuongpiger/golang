package parser

import (
	"errors"
)

func Parse(input string) (int, error) {
	if input == "" {
		return 0, errors.New("empty input")
	}

	// Imagine parsing logic here
	return 42, nil
}

