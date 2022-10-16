package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	goRoutines := 5
	wait.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func(id int) {
			defer wait.Done()
			fmt.Printf("GoroutineID %v, Hello, World!\n", id)
		}(i)
	}
	wait.Wait()
}
