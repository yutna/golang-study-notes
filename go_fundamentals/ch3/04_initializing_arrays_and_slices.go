package main

import "fmt"

func main() {
	a := [5]int{}
	b := []int{}
	c := [3]int{1, 2, 3}
	d := []int{1, 2, 3}

	var e [5]int
	var f []int

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
