package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type (
	Subscriber interface {
		Notify(interface{}) error
		Close()
	}

	Publisher interface {
		start()
		AddSubcriberCh() chan<- Subscriber
		RemoveSubcriberCh() chan<- Subscriber
		PublishingCh() chan<- interface{}
		Stop()
	}

	writerSubscriber struct {
		in     chan interface{}
		id     int
		Writer io.Writer
	}

	publisher struct {
		subscribers []Subscriber
		addSubCh    chan Subscriber
		removeSubCh chan Subscriber
		in          chan interface{}
		stop        chan struct{}
	}

	mockWriter struct {
		testingFunc func(string)
	}

	mockSubscriber struct {
		notifyTestingFunc func(interface{})
		closeTestingFunc  func()
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

func NewWriterSubscriber(id int, writer io.Writer) Subscriber {
	if writer == nil {
		writer = os.Stdout
	}

	s := &writerSubscriber{
		id:     id,
		in:     make(chan interface{}),
		Writer: writer,
	}

	go func() {
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d): %v\n", s.id, msg)
		}
	}()

	return s
}

func NewPublisher() Publisher {
	return &publisher{}
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.testingFunc(string(p))
	return len(p), nil
}

func (m *mockSubscriber) Notify(msg interface{}) error {
	m.notifyTestingFunc(msg)
	return nil
}

func (m *mockSubscriber) Close() {
	m.closeTestingFunc()
}

func (p *publisher) AddSubcriberCh() chan<- Subscriber {
	return p.addSubCh
}

func (p *publisher) RemoveSubcriberCh() chan<- Subscriber {
	return p.removeSubCh
}

func (p *publisher) PublishingCh() chan<- interface{} {
	return p.in
}

func (p *publisher) start() {
	for {
		select {
		case msg := <-p.in:
			for _, sub := range p.subscribers {
				sub.Notify(msg)
			}
		case sub := <-p.addSubCh:
			p.subscribers = append(p.subscribers, sub)
		case sub := <-p.removeSubCh:
			for i, candidate := range p.subscribers {
				if candidate == sub {
					p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
					candidate.Close()
					break
				}
			}
		case <-p.stop:
			for _, sub := range p.subscribers {
				sub.Close()
			}
			close(p.addSubCh)
			close(p.in)
			close(p.removeSubCh)

			return
		}
	}
}

func (p *publisher) Stop() {
	close(p.stop)
}

func main() {
	msg := "Hello"

	p := NewPublisher()
	var wg sync.WaitGroup
	sub := &mockSubscriber{
		notifyTestingFunc: func(msg interface{}) {
			defer wg.Done()
			s, ok := msg.(string)
			if !ok {
				fmt.Println("Invalid data type")
			}

			if s != msg {
				fmt.Println("Invalid message")
			}
		},
		closeTestingFunc: func() {
			defer wg.Done()
		},
	}

	p.AddSubcriberCh() <- sub
	p.PublishingCh() <- msg

	pubCon := p.(*publisher)
	if len(pubCon.subscribers) != 1 {
		fmt.Println("Invalid subscriber count")
	}
	p.RemoveSubcriberCh() <- sub

	if len(pubCon.subscribers) != 0 {
		fmt.Println("Invalid subscriber count")
	}
	p.Stop()
}
