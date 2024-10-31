package main

import (
	"fmt"
	"strings"
)

func main() {
	counts := map[string]int{}
	sentence := "The quick brown fox jumps over the lazy dog"
	words := strings.Fields(strings.ToLower(sentence))

	for _, w := range words {
		counts[w]++
	}

	fmt.Println(counts)
}
