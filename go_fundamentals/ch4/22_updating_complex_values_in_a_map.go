package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	data := map[int]User{}
	user1 := User{ID: 1, Name: "Kurt"}

	data[1] = user1

	user2 := data[1]
	user2.Name = "Janis"

	data[1] = user2

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", user1)
	fmt.Printf("%+v\n", user2)
}
