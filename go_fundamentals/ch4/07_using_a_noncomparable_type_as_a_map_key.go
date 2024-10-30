package main

import "fmt"

func main() {
	m := map[func()]string{}
	fmt.Println(m)
}
