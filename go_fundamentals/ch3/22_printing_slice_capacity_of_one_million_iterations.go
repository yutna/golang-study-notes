package main

import "fmt"

func main() {
	var s1 []int
	hat := cap(s1)

	for i := 0; i < 1_000_000; i++ {
		s1 = append(s1, i)
		c := cap(s1)

		if c != hat {
			fmt.Println(hat, c)
		}

		hat = c
	}
}
