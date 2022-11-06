package main

import "fmt"

type (
	Observer interface {
		update(string)
		getID() string
	}

	Item struct {
		observerList []Observer
		name         string
		inStock      bool
	}

	Customer struct {
		id string
	}
)

// Item's collection of methods
func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (s *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", s.name)
	s.inStock = true
	s.notifyAll()
}
func (s *Item) register(o Observer) {
	s.observerList = append(s.observerList, o)
}

func (s *Item) deregister(o Observer) {
	s.observerList = removeFromslice(s.observerList, o)
}

func (s *Item) notifyAll() {
	for _, observer := range s.observerList {
		observer.update(s.name)
	}
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// Customer's collection of methods
func (s *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", s.id, itemName)
}

func (s *Customer) getID() string {
	return s.id
}

func main() {
	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
