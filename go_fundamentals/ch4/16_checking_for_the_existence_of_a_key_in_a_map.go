package main

import (
	"fmt"
	"os"
)

func main() {
	users := map[string]string{
		"Kurt":  "kurt@example.com",
		"Janis": "janis@example.com",
		"Jimi":  "jimi@example.com",
		"Amy":   "amy@example.com",
	}

	key := "Kurt"

	email, ok := users[key]
	if !ok {
		fmt.Printf("Key not found: %q\n", key)
		os.Exit(1)
	}

	fmt.Printf("Found key %q: %q", key, email)
}
