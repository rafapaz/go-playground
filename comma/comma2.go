package main

import (
	"fmt"
	"bytes"
)

func main() {
	fmt.Println(comma("1234567890"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer

	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if count == 3 {
			buf.WriteByte(',')
			count = 0
		}
		buf.WriteByte(s[i])
		count++
	}

	inverted := buf.String()
	buf.Reset()

	for i := len(inverted) - 1; i >= 0; i-- {
		buf.WriteByte(inverted[i])
	}

	return buf.String()
}
