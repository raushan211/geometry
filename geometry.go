package main

import "fmt"

func main() {
	fmt.Println("enter shape: ")
	// 1 for rectangle. 2 for circle, 3 for square
	var number int
	fmt.Scanln(&number)
	fmt.Println("enter operation: ")
	// area or perimeter
	var operation string
	fmt.Scanln(&operation)

	var result int
	if number == "1" && operation == "area" {
		fmt.Println("enter lenght: ")
		var lenght int
		fmt.Scanln(&lenght)
		fmt.Println("enter breadth: ")
		var breadth int
		fmt.Scanln(&breadth)
		result = length * breadth
	}
	else if number == "1" && operation == "perimeter" {
		fmt.Println("enter lenght: ")
		var lenght int
		fmt.Scanln(&lenght)
		fmt.Println("enter breadth: ")
		var breadth int
		fmt.Scanln(&breadth)
		result = 2 * (lenght + breadth)
	} else if number == "2" && operatin == "area" {
		fmt.Println("enter radius: ")
		var radius int
		fmt.Scanln(&radius)
		result = 3.14 * radius * radius
	} else if number == "2" && operatin == "perimeter" {
		fmt.Println("enter radius: ")
		var radius int
		fmt.Scanln(&radius)
		result = 2 * 3.14 * radius 
    } else if number == "3" && operation == "area" {
		fmt.Println("enter side: ")
		var side int
		fmt.Scanln(&side)
		result = side * side
	} else if number == "3" && opearation == "perimeter" {
		fmt.Println("enter side: ")
		var side int
		fmt.Scanln(&side)
		result = 4 * side
	} else {
		fmt.Println("wrong input")
	}

}
