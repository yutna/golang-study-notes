package main

import "fmt"

func main() {
	a := make([]string, 1, 3)

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
