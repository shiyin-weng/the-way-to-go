package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	arr := []byte{'1', '2'}
	//data := []byte{'1', '2', '1', '2'}

	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	for i := 0; i < 100; i++ {
		arr = append(arr, 'a')
	}
	fmt.Println(arr)

	fmt.Println(len(arr))
	fmt.Println(cap(arr))
}

func Append(slice, data []byte) []byte {
	return append(slice, data...)
}
