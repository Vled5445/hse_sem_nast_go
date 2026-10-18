package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var num int
	fmt.Scan(&num)

	max := num
	min := num

	for i := 1; i < n; i++ {
		fmt.Scan(&num)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}
