package main

import "fmt"

type LatLong struct {
	Lat  float64
	Long float64
}

type Animal interface {
	GetLocation() LatLong
	SetLocation(LatLong)
	CanFly() bool
	Speak() string
	GetName() string
}

// The Lion Family
//

type Lion struct {
	name       string
	maneLength int
	location   LatLong
}

func (lion *Lion) GetLocation() LatLong {
	return lion.location
}

func (lion *Lion) SetLocation(loc LatLong) {
	lion.location = loc
}

func (lion *Lion) CanFly() bool {
	return false
}

func (lion *Lion) Speak() string {
	return "roar"
}

func (lion *Lion) GetManeLength() int {
	return lion.maneLength
}

func (lion *Lion) GetName() string {
	return lion.name
}

// The Pigeon Family
//

type Pigeon2 struct {
	name     string
	location LatLong
}

func (p *Pigeon2) GetLocation() LatLong {
	return p.location
}

func (p *Pigeon2) SetLocation(loc LatLong) {
	p.location = loc
}

func (p *Pigeon2) CanFly() bool {
	return false
}

func (p *Pigeon2) Speak() string {
	return "hoot"
}

func (p *Pigeon2) GetName() string {
	return p.name
}

// The symphony
// makeThemSing demonstraces how client code can work with interfaces and not worry about struct specifics
func makeThemSing(animals []Animal) {
	for _, animal := range animals {
		fmt.Println(animal.GetName() + " says " + animal.Speak())
	}
}

func main() {
	var myZoo []Animal

	Leo := Lion{
		"Leo",
		10,
		LatLong{10.40, 11.5},
	}
	myZoo = append(myZoo, &Leo)

	Tweety := Pigeon2{
		"Tweety",
		LatLong{10.40, 11.5},
	}
	myZoo = append(myZoo, &Tweety)

	makeThemSing(myZoo)

	var aAnimal Animal

	aAnimal = &Lion{
		"Leo",
		10,
		LatLong{10.40, 11.5},
	}

	fmt.Println(aAnimal)

}
