package main

import "fmt"

func main() {
	buf := []byte{'1', '2', '3'}

	firstN, rest := buf[:2], buf[2:]

	fmt.Println(firstN)
	fmt.Println(rest)
}
