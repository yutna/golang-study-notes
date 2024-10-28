package main

import "fmt"

func main() {
	var names []string

	names = append(names, "Kris")
	fmt.Println(names)

	names = append(names, "Janis", "Jimi")
	fmt.Println(names)

	users := []string{"Kris"}
	users = append(users, "Janis", "Jimi")
	fmt.Println(users)
}
