package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("1234.567890"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var dec string

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		dec = s[dot:]
		s = s[:dot]
	}

	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:] + dec
}
