package main

import "fmt"

type (
	// IPizza is the interface for the pizza
	IPizza interface {
		getPrice() int
	}

	// VeggeMania is the concrete component
	VeggeMania struct{}

	// TomatoTopping is the concrete decorator
	TomatoTopping struct {
		pizza IPizza
	}

	// CheeseTopping is the concrete decorator
	CheeseTopping struct {
		pizza IPizza
	}
)

// VeggeMania's collection of methods
func (s *VeggeMania) getPrice() int {
	return 15
}

// TomatoTopping's collection of methods
func (s *TomatoTopping) getPrice() int {
	pizzaPrice := s.pizza.getPrice()
	return pizzaPrice + 7
}

// CheeseTopping's collection of methods
func (s *CheeseTopping) getPrice() int {
	pizzaPrice := s.pizza.getPrice()
	return pizzaPrice + 10
}

// main function
func main() {
	pizza := new(VeggeMania)
	pizzaWithCheese := &CheeseTopping{pizza}
	pizzaWithCheeseAndTomato := &TomatoTopping{pizzaWithCheese}

	fmt.Println("Price of VeggeMania with cheese and tomato topping is: ", pizzaWithCheeseAndTomato.getPrice())
}
