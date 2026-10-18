package main

import "fmt"

func main() {
	var a, b float64
	var op string

	fmt.Scan(&a, &op, &b)

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Error: division by zero")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Unknown operation")
	}
}
