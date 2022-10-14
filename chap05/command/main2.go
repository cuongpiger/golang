package main

import (
	"fmt"
	"time"
)

type Command2 interface {
	Info() string
}

type ChainLogger interface {
	Next(Command2)
}

type TimePassed2 struct {
	start time.Time
}

func (t *TimePassed2) Info() string {
	return time.Since(t.start).String()
}

type Logger struct {
	NextChain ChainLogger
}

func (f *Logger) Next(c Command2) {
	time.Sleep(time.Second)

	fmt.Printf("Elapsed time from creation: %s\n", c.Info())

	if f.NextChain != nil {
		f.NextChain.Next(c)
	}
}

func main() {
	second := new(Logger)
	first := Logger{NextChain: second}

	command := &TimePassed2{start: time.Now()}

	first.Next(command)
}
