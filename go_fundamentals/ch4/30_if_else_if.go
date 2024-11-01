package main

import "fmt"

func main() {
	greet := true
	name := "Janis"

	if greet && name != "" {
		fmt.Println("Hello", name)
	} else if greet {
		fmt.Println("Hello")
	} else {
		fmt.Println("Goodbye")
	}
}
