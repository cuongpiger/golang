package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	Request struct {
		Data    interface{}
		Handler RequestHandler
	}

	RequestHandler func(interface{})
	WorkerLauncher interface {
		LaunchWorker(in chan Request)
	}
	Dispatcher interface {
		LaunchWorker(w WorkerLauncher)
		MakeRequest(r Request)
		Stop()
	}
	PrefixSuffixWorker struct {
		id      int
		prefixS string
		suffixS string
	}

	dispatcher struct {
		inCh chan Request
	}
)

func NewStringRequest(w string, id int, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data: "Hello",
		Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				fmt.Println("Invalid data type")
			}
			fmt.Println(fmt.Sprintf(w, id) + " " + s)
		},
	}

	return myRequest
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inCh <- r:
	case <-time.After(time.Second * 5):
		return
	}
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}

func NewDispatcher(b int) Dispatcher {
	return &dispatcher{
		inCh: make(chan Request, b),
	}
}

func (w *PrefixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PrefixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}

			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()

	return out
}

func (w *PrefixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}

			msg.Data = fmt.Sprintf("%s%s", s, w.suffixS)
			out <- msg
		}
		close(out)
	}()

	return out
}

func (w *PrefixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}

			msg.Handler(fmt.Sprintf("%s%s", w.prefixS, s))
		}
	}()
}

func main() {
	bufferSize := 100
	var dispatcher = NewDispatcher(bufferSize)
	workers := 3
	for i := 0; i < workers; i++ {
		worker := &PrefixSuffixWorker{
			id:      i,
			prefixS: fmt.Sprintf("WorkerID %d -> ", i),
			suffixS: " World",
		}
		dispatcher.LaunchWorker(worker)
	}

	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		req := NewStringRequest("MsgID %d -> Hello", i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()
	wg.Wait()
}
