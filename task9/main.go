package main

import "fmt"

func main() {
	var balance float64
	fmt.Scan(&balance)

	var action int
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Scan(&action)

	switch action {
	case 1:
		fmt.Println("Balance:", balance)
	case 2:
		var amount float64
		fmt.Scan(&amount)
		balance += amount
		fmt.Println("Balance:", balance)
	case 3:
		var amount float64
		fmt.Scan(&amount)
		if amount > balance {
			fmt.Println("Insufficient funds")
		} else {
			balance -= amount
			fmt.Println("Balance:", balance)
		}
	default:
		fmt.Println("Invalid action")
	}
}
