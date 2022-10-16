package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan string)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		channel <- "Hello World!"
		fmt.Println("Finished go-routine")
	}()

	message := <-channel
	fmt.Println(message)
	waitGroup.Wait()
}
