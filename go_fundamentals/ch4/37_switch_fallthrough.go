package main

import "fmt"

func RecommendActivity(temp int) {
	fmt.Printf("It is %d degrees out. You could", temp)

	switch {
	case temp <= 32:
		fmt.Print(" go icd skating.")
		fallthrough
	case temp >= 45 && temp < 90:
		fmt.Print(" go jogging.")
		fallthrough
	case temp >= 80:
		fmt.Print(" go swimming.")
		fallthrough
	default:
		fmt.Print(" or just stay home.\n")
	}
}

func main() {
	RecommendActivity(19)
	RecommendActivity(45)
	RecommendActivity(90)
}
