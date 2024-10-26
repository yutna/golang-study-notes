package main

func main() {
	strings := [4]string{"one", "two", "three", "four"}
	strings[0] = 5

	ints := []int{1, 2, 3, 4}
	ints[0] = "five"

	mixed := []string{"one", 2, "three", "four"}
}
