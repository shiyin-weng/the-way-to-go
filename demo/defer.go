package main

import "fmt"

// common usage
// open a file
// defer file.Close()

// mu.Lock()
// defer mu.Unlock()

// printHeader()
// defer printFooter()

// open a database connection
// defer disconnectFromDB()

func main() {
	//a()
	f()
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
