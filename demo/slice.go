package main

import "fmt"

func main() {
	//var arr = [5]int{0, 1, 2, 3, 4}
	//arr := make([]int, 10)
	//fmt.Println(arr)
	//s := make([]byte, 5)
	//fmt.Println(len(s))
	//fmt.Println(cap(s))

	//s1 := s[2:4]
	//fmt.Println(len(s1))
	//fmt.Println(cap(s1))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]

	fmt.Println(s1) // {'p', 'o', 'e', 'm'}
	fmt.Println(s2) // {'e', 'm'}

	s2[1] = 't'

	fmt.Println(s1) // {'p', 'o', 'e', 't'}
	fmt.Println(s2) // {'e', 't'}
}
