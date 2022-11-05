package main

import "fmt"

type (
	Originator struct {
		state string
	}

	Memento struct {
		state string
	}

	CareTaker struct {
		mementoArray []*Memento
	}
)

// Originator's collection of methods
func (s *Originator) createMemento() *Memento {
	return &Memento{s.state}
}

func (s *Originator) restoreMemento(m *Memento) {
	s.state = m.getSavedState()
}

func (s *Originator) setState(state string) {
	s.state = state
}

func (s *Originator) getState() string {
	return s.state
}

// Memento's collection of methods
func (s *Memento) getSavedState() string {
	return s.state
}

// CareTaker's collection of methods
func (s *CareTaker) addMemento(m *Memento) {
	s.mementoArray = append(s.mementoArray, m)
}

func (s *CareTaker) getMemento(index int) *Memento {
	return s.mementoArray[index]
}

// main function
func main() {
	careTaker := &CareTaker{
		mementoArray: make([]*Memento, 0),
	}
	originator := &Originator{
		state: "A",
	}

	fmt.Printf("Originator Current State: %s\n", originator.getState())
	careTaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	careTaker.addMemento(originator.createMemento())

	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	careTaker.addMemento(originator.createMemento())

	originator.restoreMemento(careTaker.getMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(careTaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
