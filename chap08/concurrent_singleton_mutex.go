package main

import (
	"fmt"
	"sync"
	"time"
)

type singleton1 struct {
	count int
	sync.RWMutex
}

var instance1 singleton1

func GetInstance1() *singleton1 {
	return &instance1
}

func (s *singleton1) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *singleton1) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}

func main() {
	singleton := GetInstance1()
	singleton2 := GetInstance1()
	n := 5_000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("before loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Printf("before loop, current count is %d\n", singleton.GetCount())
}
