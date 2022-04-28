package main

import "fmt"

var naturals chan int
var squares chan int

func main() {
	naturals = make(chan int)
	squares = make(chan int)

	go counter()
	go squarer()
	for {
		fmt.Println(<-squares)
	}
}

func counter() {
	for x := 0; ; x++ {
		naturals <- x
	}
}

func squarer() {
	for {
		x := <-naturals
		squares <- x * x
	}
}
