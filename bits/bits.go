package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var diff [32]byte
	var count int

	c1 := sha256.Sum256([]byte("a"))
	c2 := sha256.Sum256([]byte("b"))
	fmt.Printf("%b\n", c1)
	fmt.Printf("%b\n", c2)

	for i := 0; i < len(c1); i++ {
		diff[i] = c1[i] ^ c2[i]
	}

	fmt.Printf("%b\n", diff)
	
	for i := 0; i < len(diff); i++ {
		n := diff[i]
		for n > 0 {
			count += int(n & 1)
			n >>= 1
		}
	}

	fmt.Printf("We have %d different bits\n", count)

}
