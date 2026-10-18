package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var num int
	fmt.Scan(&num)

	sum := num
	max := num
	min := num
	even, odd := 0, 0

	if num%2 == 0 {
		even++
	} else {
		odd++
	}

	for i := 1; i < n; i++ {
		fmt.Scan(&num)
		sum += num

		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
		if num%2 == 1 || num%2 == -1 {
			odd++
		} else {
			even++
		}
	}

	avg := float64(sum) / float64(n)

	fmt.Println("Sum:", sum)
	fmt.Println("Average:", avg)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
	fmt.Println("Even:", even)
	fmt.Println("Odd:", odd)
}
