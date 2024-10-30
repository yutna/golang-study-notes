package main

import "fmt"

func main() {
	users := map[string]string{}

	users["Janis"] = "janis@example.com"
	email := users["Janis"]

	fmt.Println(email)
}
