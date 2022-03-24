package main

import (
	"fmt"
	"bufio"
	"os"
)

type dado struct {
	qtd int
	arq string
}

func countLines(f *os.File, counts map[string]*dado, filename string) {

	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := counts[input.Text()]; !ok {
			counts[input.Text()] = &dado{0, ""}
		}
		counts[input.Text()].qtd++
		counts[input.Text()].arq = filename
	}
}


func main() {
	counts := make(map[string]*dado)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}

	for line, d := range counts {
		if d.qtd > 1 {
			fmt.Printf("%d\t%s\t%s\n", d.qtd, line, d.arq)
		}
	}

}

