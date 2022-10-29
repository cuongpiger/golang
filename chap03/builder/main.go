package main

import "fmt"

type (
	// IBuilder is the interface for the builder
	IBuilder interface {
		setWindowType()
		setDoorType()
		setNumFloor()
		getHouse() House
	}

	// Director is the object that will use the builder
	Director struct {
		builder IBuilder
	}

	// House is the House object
	House struct {
		windowType string
		doorType   string
		floor      int
	}

	// NormalBuilder is a variant of House object
	NormalBuilder struct {
		windowType string
		doorType   string
		floor      int
	}

	// IglooBuilder is a variant of House object
	IglooBuilder struct {
		windowType string
		doorType   string
		floor      int
	}
)

// implement methods for IBuilder
func getBuilder(builderType string) IBuilder {
	switch builderType {
	case "normal":
		return newNormalBuilder()
	case "igloo":
		return newIglooBuilder()
	}

	return nil
}

// implement methods for NormalBuilder
func newNormalBuilder() *NormalBuilder {
	return new(NormalBuilder)
}

func (s *NormalBuilder) setWindowType() {
	s.windowType = "Wooden Window"
}

func (s *NormalBuilder) setDoorType() {
	s.doorType = "Wooden Door"
}

func (s *NormalBuilder) setNumFloor() {
	s.floor = 2
}

func (s *NormalBuilder) getHouse() House {
	return House{
		doorType:   s.doorType,
		windowType: s.windowType,
		floor:      s.floor,
	}
}

// implement methods for IglooBuilder object
func newIglooBuilder() *IglooBuilder {
	return new(IglooBuilder)
}

func (s *IglooBuilder) setWindowType() {
	s.windowType = "Snow Window"
}

func (s *IglooBuilder) setDoorType() {
	s.doorType = "Snow Door"
}

func (s *IglooBuilder) setNumFloor() {
	s.floor = 1
}

func (s *IglooBuilder) getHouse() House {
	return House{
		doorType:   s.doorType,
		windowType: s.windowType,
		floor:      s.floor,
	}
}

// implement methods for Director object
func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (s *Director) setBuilder(b IBuilder) {
	s.builder = b
}

func (s *Director) buildHouse() House {
	s.builder.setDoorType()
	s.builder.setWindowType()
	s.builder.setNumFloor()

	return s.builder.getHouse()
}

// main function
func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}
