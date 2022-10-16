package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type (
	GameState1 interface {
		executeState(*GameContext1) bool
	}
)

type (
	GameContext1 struct {
		SecretNumber int
		Retries      int
		Won          bool
		Next         GameState1
	}

	StartState1  struct{}
	FinishState1 struct{}
	AskState1    struct{}
)

func (s *StartState1) executeState(c *GameContext1) bool {
	c.Next = &AskState1{}

	rand.Seed(time.Now().UnixNano())
	c.SecretNumber = rand.Intn(10)
	fmt.Println("Introduce a number a number of retries to set the difficulty:")
	fmt.Fscanf(os.Stdin, "%d\n", &c.Retries)

	return true
}

func (a *AskState1) executeState(c *GameContext1) bool {
	fmt.Printf("Introduce a number between 0 and 10, you have %d tries left\n", c.Retries)

	var n int
	fmt.Fscanf(os.Stdin, "%d", &n)
	c.Retries = c.Retries - 1

	if n == c.SecretNumber {
		c.Won = true
		c.Next = &FinishState1{}
	}

	if c.Retries == 0 {
		c.Next = &FinishState1{}
	}

	return true
}

func (f *FinishState1) executeState(c *GameContext1) bool {
	if c.Won {
		println("Congrats, you won")
	} else {
		fmt.Printf("You loose. The correct number was: %d\n", c.SecretNumber)
	}

	return false
}

func main() {
	start := StartState1{}
	game := GameContext1{
		Next: &start,
	}

	for game.Next.executeState(&game) {
	}
}
