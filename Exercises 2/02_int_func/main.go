package main

import "fmt"

func main() {
	mathy := func(i int) (int, bool) {
		fmt.Println("Enter an integer: ")
		fmt.Scanln(&i)
		return i / 2, i%2 == 0
	}
	fmt.Println(mathy(0))
}

// func that takes an int and then divides by 2
// also uses modulo to check if number is even and return a bool
