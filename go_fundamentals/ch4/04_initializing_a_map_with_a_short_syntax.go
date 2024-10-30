package main

import "fmt"

func main() {
	users := map[string]string{
		"kurt@example.com":  "Kurt",
		"janis@example.com": "Janis",
		"jimi@example.com":  "Jimi",
	}

	fmt.Println(users)
}
