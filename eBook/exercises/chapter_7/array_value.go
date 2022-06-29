package main

import "fmt"

func main() {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	arr2 := arr1
	arr2[2] = 100

	fmt.Println(&arr1)
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array arr1 at index %d is %d\n", i, arr1[i])
	}
	fmt.Println()
	fmt.Println(&arr2)
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Array arr2 at index %d is %d\n", i, arr2[i])
	}
}
