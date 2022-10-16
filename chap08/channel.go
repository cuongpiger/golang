package main

func main() {
	channel := make(chan string)
	go func() {
		channel <- "Hello World!"
	}()

	message := <-channel
	println(message)
}