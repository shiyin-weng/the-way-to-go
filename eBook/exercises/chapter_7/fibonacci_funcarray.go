package main

import "fmt"

var term = 1

func main() {
	result := fibarray(term)
	for ix, fib := range result {
		fmt.Printf("The %d-th Fibonacci number is: %d\n", ix, fib)
	}
}

func fibarray(term int) []int {
	if term == 1 {
		return []int{1}
	}

	farr := make([]int, term)
	farr[0], farr[1] = 1, 1

	for i := 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}
	return farr
}
