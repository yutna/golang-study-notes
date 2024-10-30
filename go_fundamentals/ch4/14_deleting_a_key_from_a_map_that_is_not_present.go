package main

import "fmt"

func main() {
	users := map[string]string{
		"Kurt":  "kurt@example.com",
		"Janis": "janis@example.com",
		"Jimi":  "jimi@example.com",
		"amy":   "amy@example.com",
	}

	delete(users, "Unknown")

	fmt.Println(users)
}
