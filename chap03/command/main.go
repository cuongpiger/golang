package main

import "fmt"

type (
	// Command is an interface for executing a command.
	Command interface {
		execute()
	}

	// Device is the receiver interface
	Device interface {
		on()
		off()
	}

	// Button is invoker
	Button struct {
		command Command
	}

	// OnCommand is a concrete command
	OnCommand struct {
		device Device
	}

	// OffCommand is a concrete command
	OffCommand struct {
		device Device
	}

	// Tv is concrete receiver
	Tv struct {
		isRunning bool
	}
)

// Button's collection of methods
func (s *Button) press() {
	s.command.execute()
}

// OnCommand's collection of methods
func (s *OnCommand) execute() {
	s.device.on()
}

// OffCommand's collection of methods
func (s *OffCommand) execute() {
	s.device.off()
}

// Tv's collection of methods
func (s *Tv) on() {
	s.isRunning = true
	fmt.Println("TV is on")
}

func (s *Tv) off() {
	s.isRunning = false
	fmt.Println("TV is off")
}

// main function
func main() {
	tv := new(Tv)
	onCommand := &OnCommand{tv}
	offCommand := &OffCommand{tv}

	onButton := &Button{onCommand}
	onButton.press()

	offButton := &Button{offCommand}
	offButton.press()
}
