package main

import (
	"fmt"
)

func main() {
	var small int32
	var large int32
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&small)
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&large)
	fmt.Println(large % small)
}
