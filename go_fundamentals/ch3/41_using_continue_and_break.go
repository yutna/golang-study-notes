package main

import "fmt"

func main() {
	var i int

	for {
		i++

		if i == 3 {
			continue
		}

		if i == 10 {
			break
		}

		fmt.Println(i)
	}

	fmt.Println("finished")
}
