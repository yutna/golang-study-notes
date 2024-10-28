package main

import "fmt"

func main() {
	a1 := [2]string{"one", "two"}
	a2 := [2]string{}

	a2 = a1

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)

	a1[0] = "three"

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
}
