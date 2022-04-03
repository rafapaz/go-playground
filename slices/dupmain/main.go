package main

import (
	"dupmain/duplicate"
	"fmt"
)

func main() {
	s := []string{"aaa", "bbb", "bbb", "ccc", "ddd", "eee", "eee", "eee", "fff", "ggg", "hhh", "iii", "iii"}
	fmt.Printf("%s\n", s)
	s = duplicate.RemoveDup(s)
	fmt.Printf("%d - %s\n", len(s), s)
}
