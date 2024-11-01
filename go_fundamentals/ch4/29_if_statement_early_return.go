package main

import "fmt"

func main() {
	greet := true

	if greet {
		fmt.Println("Hello")
		return
	}

	fmt.Println("Goodbye")
}
