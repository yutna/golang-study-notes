package main

import "fmt"

func main() {
	users := map[string]string{}

	users["Kurt"] = "kurt@example.com"
	users["Janis"] = "janis@example.com"
	users["Jimi"] = "jimi@example.com"
	users["Amy"] = "amy@example.com"

	fmt.Println("Map length:", len(users))
}
