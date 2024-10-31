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
	data[1].Name = "Janis"

	fmt.Printf("%+v\n", data)
}
