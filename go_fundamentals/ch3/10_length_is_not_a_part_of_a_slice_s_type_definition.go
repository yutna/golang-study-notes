package main

import "fmt"

func main() {
	s1 := []string{"one", "two"}
	var s2 []string
	s3 := []int{}

	s2 = s1
	fmt.Println(s2)

	s3 = s2
}
