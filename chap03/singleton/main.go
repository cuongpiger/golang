package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var singleInstance *single

type (
	// single is a singleton
	single struct{}
)

// getInstance returns the singleton instance
func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating Single Instance Now")
			singleInstance = new(single)
		} else {
			fmt.Println("Single Instance already created")
		}
	} else {
		fmt.Println("Single Instance already created")
	}

	return singleInstance
}

// main function
func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.
	fmt.Scanln()
}