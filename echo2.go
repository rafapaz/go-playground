package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "",""
	_, arg := range os.Args[1:]
	fmt.Println(arg)
}
