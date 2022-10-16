package main

import "time"

func main() {
	channel := make(chan string, 1)
	go func(ch chan<- string) {
		channel <- "Hello World!"
		println("Finished go-routine")
	}(channel)
	time.Sleep(1 * time.Second)

	message := <-channel
	println(message)
}
