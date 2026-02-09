package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	var op string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&op)

	var result int
	if op == "+" {
		result = num1 + num2
	} else if op == "-" {
		result = num1 - num2
	} else if op == "*" {
		result = num1 * num2
	} else if op == "/" {
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Division by zero")
			return
		}
	} else {
		fmt.Println("Invalid operator")
		return
	}
	fmt.Printf("Result: %d\n", result)
}
