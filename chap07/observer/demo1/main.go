package main

import "fmt"

type (
	Observer interface {
		Notify(string)
	}

	Publisher struct {
		ObserverList []Observer
	}

	TestObserver struct {
		ID      int
		Message string
	}
)

func (s *Publisher) AddObserver(o Observer) {
	s.ObserverList = append(s.ObserverList, o)
}

func (s *Publisher) RemoveObserver(o Observer) {
	for i, v := range s.ObserverList {
		if v == o {
			s.ObserverList = append(
				s.ObserverList[:i],
				s.ObserverList[i+1:]...)
			break
		}
	}
}

func (s *Publisher) NotifyObservers(msg string) {
	fmt.Printf("Publisher received message '%s' to notify observers\n", msg)
	for _, observer := range s.ObserverList {
		observer.Notify(msg)
	}
}

func (p *TestObserver) Notify(m string) {
	fmt.Printf("Observer %d: message '%s' received \n", p.ID, m)
	p.Message = m
}
