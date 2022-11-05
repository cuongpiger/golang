package main

import "fmt"

type (
	Train interface {
		arrive()
		depart()
		permitArrival()
	}

	Mediator interface {
		canArrive(train Train) bool
		notifyAboutDeparture()
	}

	PassengerTrain struct {
		mediator Mediator
	}

	FreightTrain struct {
		mediator Mediator
	}

	StationManager struct {
		isPlatformFree bool
		trainQueue     []Train
	}
)

// PassengerTrain's collection of methods
func (s *PassengerTrain) arrive() {
	if !s.mediator.canArrive(s) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (s *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	s.mediator.notifyAboutDeparture()
}

func (s *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	s.arrive()
}

// FreightTrain's collection of methods
func (s *FreightTrain) arrive() {
	if !s.mediator.canArrive(s) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (s *FreightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	s.mediator.notifyAboutDeparture()
}

func (s *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	s.arrive()
}

// StationManager's collection of methods
func newStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}

	if len(s.trainQueue) > 0 {
		train := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		train.permitArrival()
	}
}

// main function
func main() {
	stationManager := newStationManager()
	passengerTrain := &PassengerTrain{mediator: stationManager}
	freightTrain := &FreightTrain{mediator: stationManager}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
