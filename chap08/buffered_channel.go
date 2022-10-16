package main

import "time"

func main() {
	channel := make(chan string, 1)
	go func() {
		channel <- "Hello World!"
		channel <- "Hello World! 2"
		print("Finished go-routine")
	}()

	time.Sleep(1 * time.Second)
	message := <-channel
	println(message)
	message = <-channel
	println(message)
}
