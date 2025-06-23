package main

import (
	"fmt"
	"sync"
)

type Sound struct {
	Sound []int `json:"sound"`
	mu    sync.Mutex
}

func newSound() *Sound {
	return &Sound{
		Sound: make([]int, 0),
		mu:    sync.Mutex{},
	}
}

func main() {
	repeatCount := 55

	infinite := repeatCount < 0

	for i := 0; infinite || i < repeatCount; i++ {
		fmt.Printf("Iteration %d\n", i+1)
	}

}
