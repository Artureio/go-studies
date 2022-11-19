package main

import (
	"fmt"
	clients "go-studies/bank/clients"
	bank "go-studies/bank/transactions"
)

func main() {
	artur := clients.CreateNewOwner("Artur", 21, "21919291299", "91827341892", "Developer")
	pedro := clients.CreateNewOwner("Pedro", 21, "11999999999", "12345678910", "Analyst")
	contaArtur := bank.NewAccount(artur, 1, 1, 0)
	contaPedro := bank.NewAccount(pedro, 3, 4, 0)
	contaArtur.UnlockAccount()
	contaPedro.UnlockAccount()
	fmt.Println(artur)
	fmt.Println(pedro)

	fmt.Println(contaArtur.Deposit(800))
	fmt.Println(contaPedro.Deposit(900))

	bank.Transfer(&contaArtur, &contaPedro, 300)
	fmt.Println("test")

	fmt.Println(artur)
	fmt.Println(pedro)
}
