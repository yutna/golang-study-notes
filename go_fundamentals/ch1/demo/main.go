package main

import (
	"fmt"
	"log"

	one "github.com/gofrs/uuid"
	three "github.com/gofrs/uuid/v3"
)

func main() {
	id1 := one.NewV4()
	fmt.Println(id1)

	id3, err := three.NewV4()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id3)
}
