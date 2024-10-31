package main

import "fmt"

func main() {
	months := map[int]string{
		1:  "January",
		2:  "Febuary",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	keys := make([]int, 0, len(months))

	for k := range months {
		keys = append(keys, k)
	}

	fmt.Printf("keys: %+v\n", keys)
}
