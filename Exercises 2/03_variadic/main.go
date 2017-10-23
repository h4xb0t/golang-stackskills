package main

import (
	"fmt"
)

// func max grabs an unlimited number of ints
// and returns only one int which is the largest
func max(nums ...int) int {
	var largest int
	for _, i := range nums {
		if i > largest {
			largest = i
		}
	}
	return largest
}

// sets greatest to a list of numbers and prints
func main() {
	greatest := max(10, 20, 15, 8, 9, 6)
	fmt.Println(greatest)
}
