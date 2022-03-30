package main

import (
	"fmt"
)

func main() {
	a := []int{1,2,3,4,5,6,7,8,9}
	a = rotate(a, 3)
	fmt.Println(a)
}

func rotate(s []int, i int) []int {
	bef := s[:i]
	after := s[i:]
	return append(after, bef...)
}
