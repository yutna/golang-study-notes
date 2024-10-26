package main

import "fmt"

func main() {
	a1 := [2]string{"one", "two"}
	var a2 [2]string
	a3 := [3]string{}

	a2 = a1
	fmt.Println(a2)

	a3 = a2
}
