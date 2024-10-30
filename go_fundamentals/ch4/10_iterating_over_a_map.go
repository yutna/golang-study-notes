package main

import "fmt"

func main() {
	users := map[string]string{
		"Kurt":  "kurt@example.com",
		"Janis": "janis@example.com",
		"Jimi":  "jimi@example.com",
		"Amy":   "amy@example.com",
	}

	for key, value := range users {
		fmt.Printf("%s <%s>\n", key, value)
	}
}
