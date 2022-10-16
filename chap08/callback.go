package main

import (
	"fmt"
)

func toUpperSync(word string, f func(int) int) int {
	fmt.Printf("Word: %s", word)
	a, b := 1, 2
	return f(a + b)
}

func main() {
	res := toUpperSync("hello", func(res int) int {
		fmt.Printf("Word: %d\n", res)

		return 2
	})

	fmt.Printf("Result: %d", res)
}
