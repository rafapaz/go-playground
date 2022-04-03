package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var count map[string]int
	count = make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		count[scanner.Text()]++
	}

	printSorted(count)

	//fmt.Printf("%#v\n", count)
}

func printSorted(count map[string]int) {

	for i := 0; i < len(count); i++ {
		maxv, maxk := 0, ""
		for k, v := range count {
			if v > maxv {
				maxk = k
				maxv = v
			}
		}
		fmt.Printf("%s = %d\n", maxk, maxv)
		count[maxk] = 0
	}
}
