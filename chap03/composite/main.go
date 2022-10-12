package composite

type (
	Trainer interface {
		Train()
	}

	Swimmer interface {
		Swim()
	}
)

type (
	Athlete            struct{}
	SwimmerImplementor struct{}
	CompositeSwimmerA  struct {
		MyAthlete Athlete
		MySwim    *func()
	}
	CompositeSwimmerB struct {
		Trainer
		Swimmer
	}
	Animal struct{}
	Shark  struct {
		Animal
		Swim func()
	}
	Tree struct {
		LeafValue int
		Right     *Tree
		Left      *Tree
	}
	Parent struct {
		SomeField int
	}
	Son struct {
		P Parent
	}
)

func (a *Athlete) Train() {
	println("Training")
}

func (s *SwimmerImplementor) Swim() {
	println("Swimming!")
}

func (r *Animal) Eat() {
	println("Eating")
}

func Swim() {
	println("Swimming!")
}

func GetParentField(p Parent) int {
	return p.SomeField
}
