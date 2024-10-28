package main

import "fmt"

func main() {
	names := [4]string{"Kurt", "Janis", "Jimi", "Amy"}
	slicesOnly(names[:])
}

func slicesOnly(names []string) {
	for _, name := range names {
		fmt.Println(name)
	}
}
