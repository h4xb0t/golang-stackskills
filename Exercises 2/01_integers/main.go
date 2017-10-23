package main

import "fmt"

func main() {
	fmt.Println(mathy(2))
}

// func that takes an int and then divides by 2
// also uses modulo to check if number is even and return a bool

func mathy(i int) (x int, y bool) {
	x = i / 2
	y = i%2 == 0
	return
}
