package main

import "fmt"

type Simple struct {
	ID   int
	Name string
}

func main() {
	simple := map[Simple]string{}
	sk := Simple{ID: 1, Name: "Kurt"}
	simple[sk] = "kurt@example.com"
	fmt.Println(simple)
}
