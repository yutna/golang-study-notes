package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(letters)
	fmt.Println(letters[2:5])

	fmt.Println(letters[4:len(letters)])
	fmt.Println(letters[4:])

	fmt.Println(letters[0:4])
	fmt.Println(letters[:4])
}
