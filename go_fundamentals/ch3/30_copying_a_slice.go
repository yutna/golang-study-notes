package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"Kurt", "Jenis", "Jimi", "Amy"}
	fmt.Println(names)

	subset := make([]string, 3)
	// subset := []string{}
	copy(subset, names[:3])
	fmt.Println(subset)

	for i, g := range subset {
		subset[i] = strings.ToUpper(g)
	}

	fmt.Println(subset)
	fmt.Println(names)
}
