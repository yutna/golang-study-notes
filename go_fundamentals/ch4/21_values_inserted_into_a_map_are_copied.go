package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	data := map[int]User{}
	user := User{ID: 1, Name: "Kurt"}

	data[1] = user
	user.Name = "Janis"

	fmt.Printf("User: %+v\n", user)
	fmt.Printf("Map: %+v\n", data)
}
