package main

import "fmt"

type (
	// IGun is the interface of Gun object
	IGun interface {
		setName(name string)
		setPower(power int)
		getName() string
		getPower() int
	}

	// Gun represents a Gun object
	Gun struct {
		name  string
		power int
	}

	// AK47 is a variant of Gun object which inherits from Gun object
	AK47 struct {
		Gun
	}

	// Musket is a variant of Gun object which inherits from Gun object
	Musket struct {
		Gun
	}
)

// implement methods of Gun object
func (s *Gun) setName(name string) {
	s.name = name
}

func (s *Gun) setPower(power int) {
	s.power = power
}

func (s *Gun) getName() string {
	return s.name
}

func (s *Gun) getPower() int {
	return s.power
}

// implement methods of AK47 object
func newAK47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// implement methods of Musket object
func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// implement methods of GunFactory object
func getGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return newAK47(), nil
	case "musket":
		return newMusket(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed")
}

// supplementary functions
func printDetails(g IGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}

// main function
func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}