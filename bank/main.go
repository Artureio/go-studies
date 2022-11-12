package main

import "fmt"

type bankAccount struct {
	owner   string
	agency  int
	account int
	balance float32
}

func main() {
	artur := newAccount("Artur", 123, 123456, 100)
	pedro := newAccount("Pedro", 456, 654321, 550)

	fmt.Println(artur)
	fmt.Println(pedro)

	fmt.Println(artur.deposit(80))
	fmt.Println(artur.withdraw(10))

	transfer(&artur, &pedro, 100)

	fmt.Println(artur)
	fmt.Println(pedro)
}

func newAccount(owner string, agency int, account int, balance float32) bankAccount {
	return bankAccount{owner, agency, account, balance}
}

func (account *bankAccount) withdraw(value float32) float32 {
	canWithdraw := value <= account.balance && value > 0

	if canWithdraw {
		account.balance -= value
	} else {
		fmt.Println("Insufficient funds")
	}
	return account.balance
}

func (account *bankAccount) deposit(value float32) float32 {
	if value > 0 {
		account.balance += value
	} else {
		fmt.Println("Invalid value")
	}
	return account.balance
}

func transfer(origin *bankAccount, destiny *bankAccount, value float32) {
	origin.withdraw(value)
	destiny.deposit(value)
}
