package main

import (
	"fmt"
	"sync"
	"time"
)

type (
	Subscriber interface {
		Notify(interface{}) error
		Close()
		GetID() int
	}

	Publisher interface {
		start()
		AddSubcriberCh(Subscriber)
		RemoveSubcriberCh(Subscriber)
		PublishingCh(string)
		Stop()
	}

	writerSubscriber struct {
		in chan interface{}
		id int
	}

	publisher struct {
		wg          *sync.WaitGroup
		subscribers []Subscriber
		addSubCh    chan Subscriber
		removeSubCh chan Subscriber
		in          chan interface{}
		stop        chan struct{}
	}
)

func (s *writerSubscriber) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("panic: %v", rec)
		}
	}()

	select {
	case s.in <- msg:
	case <-time.After(1 * time.Second):
		err = fmt.Errorf("timeout")
	}

	return
}

func (s *writerSubscriber) Close() {
	close(s.in)
}

func (s *writerSubscriber) GetID() int {
	return s.id
}

func NewWriterSubscriber(id int) Subscriber {
	s := &writerSubscriber{
		id: id,
		in: make(chan interface{}),
	}

	go func() {
		for msg := range s.in {
			fmt.Printf("User ID %d received message: %v\n", s.id, msg)
		}
	}()

	return s
}

func NewPublisher(wg *sync.WaitGroup) Publisher {
	return &publisher{
		wg:          wg,
		addSubCh:    make(chan Subscriber),
		removeSubCh: make(chan Subscriber),
		in:          make(chan interface{}),
		stop:        make(chan struct{}),
	}
}

func (p *publisher) AddSubcriberCh(sub Subscriber) {
	p.addSubCh <- sub
}

func (p *publisher) RemoveSubcriberCh(sub Subscriber) {
	p.removeSubCh <- sub
}

func (p *publisher) PublishingCh(msg string) {
	p.in <- msg
}

func (p *publisher) start() {
	for {
		select {
		case msg := <-p.in:
			tmp := p.subscribers
			for _, sub := range tmp {
				sub.Notify(msg)
			}
		case sub := <-p.addSubCh:
			p.wg.Add(1)
			p.subscribers = append(p.subscribers, sub)
		case sub := <-p.removeSubCh:
			for i, candidate := range p.subscribers {
				if candidate.GetID() == sub.GetID() {
					p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
					candidate.Close()
					p.wg.Done()
					break
				}
			}
		case <-p.stop:
			for _, sub := range p.subscribers {
				sub.Close()
				fmt.Println("Close subcriber", sub.GetID())
				p.wg.Done()
			}
			close(p.addSubCh)
			close(p.in)
			close(p.removeSubCh)

			return
		}
	}
}

func (p *publisher) Stop() {
	p.stop <- struct{}{}
}

func main() {
	wg := &sync.WaitGroup{}
	pub := NewPublisher(wg)
	go pub.start()

	sub1 := NewWriterSubscriber(1)
	sub2 := NewWriterSubscriber(2)
	sub3 := NewWriterSubscriber(3)

	pub.AddSubcriberCh(sub1)
	pub.AddSubcriberCh(sub2)
	pub.AddSubcriberCh(sub3)

	pub.PublishingCh("My name is Manh Cuong")
	pub.PublishingCh("hello")
	pub.RemoveSubcriberCh(sub3)
	pub.PublishingCh("world")

	pub.Stop()

	wg.Wait()
	fmt.Println("Done")
}
