package main

import "fmt"

func main() {
	users := map[string]int{
		"Kurt":  27,
		"Janis": 15,
		"Jimi":  40,
	}

	name := "Amy"

	age, ok := users[name]
	if ok {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	fmt.Printf("Could't find %s in the users map\n", name)
}
