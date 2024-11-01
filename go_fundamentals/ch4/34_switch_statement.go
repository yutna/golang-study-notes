package main

import "fmt"

func main() {
	month := 3

	switch {
	case month == 1:
		fmt.Println("January")
	case month == 2:
		fmt.Println("Febuary")
	case month == 3:
		fmt.Println("March")
	case month == 4:
		fmt.Println("April")
	case month == 5:
		fmt.Println("May")
	case month == 6:
		fmt.Println("June")
	case month == 7:
		fmt.Println("July")
	case month == 8:
		fmt.Println("August")
	case month == 9:
		fmt.Println("September")
	case month == 10:
		fmt.Println("October")
	case month == 11:
		fmt.Println("November")
	case month == 12:
		fmt.Println("December")
	}
}
