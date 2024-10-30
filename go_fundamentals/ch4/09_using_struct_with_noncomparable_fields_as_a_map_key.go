package main

import "fmt"

type Complex struct {
	Data map[string]string
	Fn   func() error
}

func main() {
	complex := map[Complex]string{}
	fmt.Println(complex)
}
