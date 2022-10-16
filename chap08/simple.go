package main

import "time"

func helloWorld() {
	println("Hello, World!")
}

func main() {
	go helloWorld()
	time.Sleep(1 * time.Second)
}
