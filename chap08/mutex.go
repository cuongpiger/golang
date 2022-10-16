package main

import (
	"sync"
	"time"
)

type (
	Counter struct {
		sync.Mutex
		value int
	}
)

func main() {
	counter := Counter{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			//counter.Lock()
			counter.value++
			//defer counter.Unlock()
		}(i)
	}

	counter.Lock()
	defer counter.Unlock()
	time.Sleep(1 * time.Second)

	println(counter.value)
}
