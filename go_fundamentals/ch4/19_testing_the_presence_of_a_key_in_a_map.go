package main

import "fmt"

func main() {
	words := map[string]int{
		"brown": 1,
		"dog":   1,
		"fox":   1,
		"jumps": 1,
		"lazy":  1,
		"over":  1,
		"quick": 1,
		"the":   2,
	}

	_, exists := words["foo"]

	fmt.Println(exists)
}