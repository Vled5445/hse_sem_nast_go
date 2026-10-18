package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n > 0 {
		fmt.Println("Positive")
	} else if n < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

	if n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	fmt.Println("Divisible by 3:", n%3 == 0)
	fmt.Println("Divisible by 5:", n%5 == 0)
}
