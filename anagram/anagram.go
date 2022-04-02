package main

import (
	"fmt"
	"os"
)

func main() {
	word1 := os.Args[1]
	word2 := os.Args[2]

	if len(word1) != len(word2) {
		fmt.Println("Not an anagram")
		return
	}

	for i, j := 0, len(word2)-1; i < len(word1); i, j = i+1, j-1 {
		if word1[i] != word2[j] {
			fmt.Println("Not an anagram")
			return
		}
	}

	fmt.Println("YES, it is an anagram!!!")
	return

}
