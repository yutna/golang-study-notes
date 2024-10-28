package main

import "fmt"

func main() {
	var names []string
	names = append(names, "Kris")
	fmt.Println(names)

	more := []string{"Janis", "Jimi"}
	names = append(names, more...)
	fmt.Println(names)
}
