package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	v, ok := counts["dddd"]
	fmt.Printf("v = %d\n", v)
	fmt.Printf("OK = %b\n", ok)

	v, ok = counts["teste"]
	fmt.Printf("v = %d\n", v)
	fmt.Printf("OK = %b\n", ok)


	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

