package main

import (
	"fmt"

	"./rotate"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a = rotate.Rotate(a, 3)
	fmt.Println(a)
}
