package main

import "fmt"

type (
	// ISportsFactory is the abstract factory interface for the factory
	ISportsFactory interface {
		makeShoe() IShoe
		makeShirt() IShirt
	}

	// IShoe is the abstract product
	IShoe interface {
		setLogo(logo string)
		setSize(size int)
		getLogo() string
		getSize() int
	}

	// IShirt is the abstract product
	IShirt interface {
		setLogo(logo string)
		setSize(size int)
		getLogo() string
		getSize() int
	}

	// Shoe is the concrete product
	Shoe struct {
		logo string
		size int
	}

	// Shirt is the concrete product
	Shirt struct {
		logo string
		size int
	}

	// Adidas is the concrete factory
	Adidas struct{}

	// AdidasShoe is the concrete product
	AdidasShoe struct {
		Shoe
	}

	// AdidasShirt is the concrete product
	AdidasShirt struct {
		Shirt
	}

	// Nike is the concrete factory
	Nike struct{}

	// NikeShoe is the concrete product
	NikeShoe struct {
		Shoe
	}

	// NikeShirt is the concrete product
	NikeShirt struct {
		Shirt
	}
)

// ISportFactory collection of methods
func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return new(Adidas), nil
	case "nike":
		return new(Nike), nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

// Shoe collection of methods
func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

// Shirt collection of methods
func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() int {
	return s.size
}

// Adidas' collection of methods
func (s *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (s *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

// Nike's collection of methods
func (s *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (s *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 14,
		},
	}
}

// supplenmentary methods
var printShoeDetails = func(s IShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

var printShirtDetails = func(s IShirt) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

// main function
func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()
	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()
	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}
