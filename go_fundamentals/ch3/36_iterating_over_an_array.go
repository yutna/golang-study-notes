package main

import "fmt"

func main() {
	names := [4]string{"Kurt", "Janis", "Jimi", "Amy"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
