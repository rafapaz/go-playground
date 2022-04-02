package main

import (
	"dupmain/duplicate"
	"fmt"
)

func main() {
	s := []string{"aaa", "bbb", "bbb", "ccc", "ddd", "eee", "eee", "eee", "fff", "ggg", "hhh", "iii", "iii"}
	fmt.Println(s)
	duplicate.RemoveDup(s)
	fmt.Println(s)
}
