package main

import "fmt"

func main() {
	data := map[int]string{}
	data[1] = "Hello, world!"

	value := data[10]

	fmt.Printf("%q", value)
}
