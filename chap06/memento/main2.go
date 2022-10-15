package main

import "fmt"

type (
	Volume byte
	Mute   bool

	Command interface {
		GetValue() interface{}
	}

	Memento struct {
		memento Command
	}

	originator1 struct {
		Command Command
	}

	careTaker1 struct {
		mementoStack []Memento
	}

	MementoFacade struct {
		originator originator1
		careTaker  careTaker1
	}
)

func (v Volume) GetValue() interface{} {
	return v
}

func (m Mute) GetValue() interface{} {
	return m
}

func (o *originator1) NewMemento() Memento {
	return Memento{memento: o.Command}
}

func (o *originator1) ExtractAndStoreState(m Memento) {
	o.Command = m.memento
}

func (c *careTaker1) Push(m Memento) {
	c.mementoStack = append(c.mementoStack, m)
}

func (c *careTaker1) Pop() Memento {
	if len(c.mementoStack) > 0 {
		memento := c.mementoStack[len(c.mementoStack)-1]
		c.mementoStack = c.mementoStack[:len(c.mementoStack)-1]
		return memento
	}

	return Memento{}
}

func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.Command = s
	m.careTaker.Push(m.originator.NewMemento())
}

func (m *MementoFacade) RestoreSettings() Command {
	m.originator.ExtractAndStoreState(m.careTaker.Pop())
	return m.originator.Command
}

func assertAndPrint(c Command) {
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}

func main() {
	m := MementoFacade{}
	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(false))

	assertAndPrint(m.RestoreSettings())
	assertAndPrint(m.RestoreSettings())
}
