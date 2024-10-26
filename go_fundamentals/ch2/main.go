package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	i, err := strconv.ParseInt("-42", 0, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-42:", i)

	o, err := strconv.ParseInt("0x2A", 0, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("0x2A:", o)

	u, err := strconv.ParseUint("42", 0, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("42 (uint):", u)

	f, err := strconv.ParseFloat("42.12345", 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("42.12345:", f)

	a, err := strconv.Atoi("42")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%[1]v, [%[1]T]\n", a)
}
