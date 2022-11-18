package main

import (
	"fmt"
	c "go-studies/bank/transactions"
)

func main() {
	artur := c.NewAccount("Artur", 123, 123456, 100)
	pedro := c.NewAccount("Pedro", 456, 654321, 550)

	fmt.Println(artur)
	fmt.Println(pedro)

	fmt.Println(artur.Deposit(80))
	fmt.Println(artur.Withdraw(10))

	c.Transfer(&artur, &pedro, 100)

	fmt.Println(artur)
	fmt.Println(pedro)
}
