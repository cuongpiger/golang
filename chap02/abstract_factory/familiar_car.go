package abstract_factory

type FamiliarCar struct{}

func (f *FamiliarCar) GetDoors() int {
	return 5
}

func (f *FamiliarCar) GetWheels() int {
	return 5
}

func (f *FamiliarCar) GetSeats() int {
	return 5
}
