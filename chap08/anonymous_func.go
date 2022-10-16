package main

import (
	"fmt"
	"time"
)

func main() {
	go func(s string) {
		fmt.Printf(s)
	}("Hello world!")

	time.Sleep(1 * time.Second)
}