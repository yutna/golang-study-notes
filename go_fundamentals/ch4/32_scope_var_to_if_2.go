package main

import "fmt"

func main() {
	users := map[string]int{
		"Kurt":  27,
		"Janis": 15,
		"Jimi":  40,
	}

	name := "Amy"

	if age, ok := users[name]; ok {
		fmt.Printf("%s is %d years old\n", name, age)
		return
	}

	fmt.Printf("Couldn't find %s in the users map\n", name)
}
