package main

import "fmt"

func main() {
	fmt.Println("enter shape: ")
	// 1 for rectangle. 2 for circle, 3 for square
	var number int
	fmt.Scanln(&number)
	fmt.Println("enter operation: ")
	// 1 for area, 2 for perimeter
	var operation int
	fmt.Scanln(&operation)

	var result float32
	if number == 1 && operation == 1 {
		fmt.Println("enter length: ")
		var length int
		fmt.Scanln(&length)
		fmt.Println("enter breadth: ")
		var breadth int
		fmt.Scanln(&breadth)
		result = length * breadth
	} else if number == 1 && operation == 2 {
		fmt.Println("enter length: ")
		var length int
		fmt.Scanln(&length)
		fmt.Println("enter breadth: ")
		var breadth int
		fmt.Scanln(&breadth)
		result = 2 * (length + breadth)
	} else if number == 2 && operation == 1 {
		fmt.Println("enter radius: ")
		var radius float32
		fmt.Scanln(&radius)
		result = 3.14 * radius * radius
	} else if number == 2 && operation == 2 {
		fmt.Println("enter radius: ")
		var radius float32
		fmt.Scanln(&radius)
		result = 2 * 3.14 * radius
	} else if number == 3 && operation == 1 {
		fmt.Println("enter side: ")
		var side float32
		fmt.Scanln(&side)
		result = side * side
	} else if number == 3 && opearation == 2 {
		fmt.Println("enter side: ")
		var side float32
		fmt.Scanln(&side)
		result = 4 * side
	} else {
		fmt.Println("wrong input")
	}
	fmt.Println(result)
}
