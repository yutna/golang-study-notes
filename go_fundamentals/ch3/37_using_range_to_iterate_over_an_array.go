package main

import "fmt"

func main() {
	names := [4]string{"Kurt", "Janis", "Jimi", "Amy"}

	for i, n := range names {
		fmt.Printf("%d %s\n", i, n)
	}
}
