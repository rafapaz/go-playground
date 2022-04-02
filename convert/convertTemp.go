package main

import (
	"fmt"
	"./tempconv"
)

func main() {
	fmt.Printf("100° C = %f° F\n", tempconv.CToF(tempconv.BoilingC))
	fmt.Printf("100° C = %f° K\n", tempconv.CToK(tempconv.BoilingC))
}
