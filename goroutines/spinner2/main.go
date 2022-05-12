package main

import (
	"fmt"
	"time"
)

// Números pequenos (até 40) funciona. Números maiores trava no final (não faz close do chanel)

func main() {
	percent := make(chan float32)
	go spinner(100*time.Millisecond, percent)
	const n = 22
	fibN := fib(n, n, percent) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration, in <-chan float32) {
	for {
		for _, r := range `-\|/` {
			perc := <-in
			fmt.Printf("\r%c %.2f%%", r, perc)
			time.Sleep(delay)
		}
	}
}

func fib(x int, orig int, out chan<- float32) int {
	perc := 100 - (float32(x)/float32(orig))*100
	if x < 2 {
		if out != nil {
			close(out)
		}
		return x
	}
	if out != nil {
		out <- perc
	}
	return fib(x-1, orig, out) + fib(x-2, orig, nil)
}
