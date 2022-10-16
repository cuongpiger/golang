package main

import "sync"

func main() {
	var wait sync.WaitGroup
	wait.Add(1)

	go func() {
		println("Hello, World!")
		wait.Add(-1)
	}()

	wait.Wait()
}