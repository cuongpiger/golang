package main

import "fmt"

func LaunchPipeline(amount int) int {
	firstCh := generator(amount)
	secondCh := power(firstCh)
	thirdCh := sum(secondCh)
	result := <-thirdCh

	//return <-sum(power(generator(amount)))
	return result
}

func generator(max int) <-chan int {
	outChInt := make(chan int, 100)
	go func() {
		for i := 0; i < max; i++ {
			outChInt <- i
		}
		close(outChInt)
	}()

	return outChInt
}

func power(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()

	return out
}

func main() {
	table := [][]int{
		{3, 14},
		{5, 55},
	}

	var res int
	for _, row := range table {
		res = LaunchPipeline(row[0])
		if res != row[1] {
			fmt.Printf("Expected %d, got %d\n", row[1], res)
		}
	}
}
