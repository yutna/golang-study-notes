package main

import "fmt"

func main() {
	var users map[string]string
	users = make(map[string]string)

	users["kurt@example.com"] = "Kurt"
	users["janis@example.com"] = "Janis"
	users["jimi@example.com"] = "Jimi"

	fmt.Println(users)
}
