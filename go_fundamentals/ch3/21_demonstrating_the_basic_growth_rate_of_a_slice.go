package main

import "fmt"

func main() {
	names := []string{}
	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))

	names = append(names, "Kurt")
	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))

	names = append(names, "Janis")
	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))

	names = append(names, "Jimi")
	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))

	names = append(names, "Amy")
	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))

	names = append(names, "Brian")
	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))
}
