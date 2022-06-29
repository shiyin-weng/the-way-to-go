package main

import (
	"fmt"
	"time"
)

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
	//f()
	fmt.Println(f1())

	start := time.Now()
	time.Sleep(time.Second)
	end := time.Now()
	fmt.Println(end.Sub(start))
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

func f1() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
