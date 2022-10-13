package main

import (
	"fmt"
	"time"
)

type Command1 interface {
	Info() string
}

type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

type HelloMessage struct{}

func (h HelloMessage) Info() string {
	return "Hello world!"
}

func main() {
	var timeCommand Command1
	timeCommand = &TimePassed{time.Now()}

	var helloCommand Command1
	helloCommand = &HelloMessage{}

	time.Sleep(time.Second)

	fmt.Println(timeCommand.Info())
	fmt.Println(helloCommand.Info())
}
