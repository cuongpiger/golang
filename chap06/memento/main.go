package main

import "fmt"

type (
	State struct {
		Description string
	}

	originator struct {
		state State
	}

	memento struct {
		state State
	}

	careTaker struct {
		mementoList []memento
	}
)

func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}

func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}

func (c *careTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Memento(i int) (memento, error) {
	if len(c.mementoList) <= i || i < 0 {
		return memento{}, fmt.Errorf("index out of range")
	}

	return c.mementoList[i], nil
}
