package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	pos, neg, zero := 0, 0, 0

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num > 0 {
			pos++
		} else if num < 0 {
			neg++
		} else {
			zero++
		}
	}

	fmt.Println("Positive:", pos)
	fmt.Println("Negative:", neg)
	fmt.Println("Zero:", zero)
}
