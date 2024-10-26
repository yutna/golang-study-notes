package main

import "fmt"

func main() {
	var a [5]int
	var b [4]string
	var c [3]bool
	var d []string

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", d) // 🤔
}
