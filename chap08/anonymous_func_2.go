package main

import "time"

func main() {
	messagePrinter := func(s string) {
		println(s)
	}

	go messagePrinter("Hello, World!")
	go messagePrinter("Hello, World 2!")
	time.Sleep(2 * time.Second)
}