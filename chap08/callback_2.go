package main

import (
	"fmt"
	"strings"
	"sync"
)

var wait sync.WaitGroup

func toUpperAsync(s string, callback func(string)) {
	go func() {
		callback(strings.ToUpper(s))
	}()
}

func main() {
	wait.Add(1)
	toUpperAsync("Hello Callback", func(v string) {
		fmt.Printf("Callback: %v\n", v)
		wait.Done()
	})

	fmt.Println("Waiting for callback...")
	wait.Wait()
}
