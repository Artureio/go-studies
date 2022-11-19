package contaCorrente

import (
	"fmt"
	"go-studies/bank/clients"
	"os"
)

type Owner clients.Owner

type BankAccount struct {
	Owner          clients.Owner
	Agency         int
	Account        int
	balance        float32
	accountBlocked bool
	FirstUse       bool
}

func (Account *BankAccount) CheckIfAccountIsBlocked() {
	if Account.accountBlocked || Account.FirstUse {
		fmt.Println(Account.Owner.Name + ", you need to unblock your Account to make this transaction")
		os.Exit(-1)
	}
}
func (Account *BankAccount) UnlockAccount() string {
	if Account.accountBlocked || Account.FirstUse {
		Account.accountBlocked = false
		Account.FirstUse = false
		return Account.Owner.Name + ", you Account has been unlocked!"
	} else {
		return Account.Owner.Name + ", Account already unlocked"
	}
}

func (Account *BankAccount) Withdraw(value float32) float32 {
	Account.CheckIfAccountIsBlocked()

	canWithdraw := value <= Account.balance && value > 0

	if canWithdraw {
		Account.balance -= value
	} else {
		fmt.Println("Insufficient funds")
	}
	return Account.balance
}

func (Account *BankAccount) Deposit(value float32) float32 {

	if value > 0 {
		Account.balance += value
	} else {
		fmt.Println("Invalid value")
	}
	return Account.balance
}

// Implement NewAccount method for Owner struct
// func (Owner *Owner) NewAccount(Agency int, Account int) BankAccount {
// 	return BankAccount{Owner, Agency, Account, balance, false, true}
// }

func NewAccount(Owner clients.Owner, Agency int, Account int, balance float32) BankAccount {
	return BankAccount{Owner, Agency, Account, balance, false, true}
}

func Transfer(origin *BankAccount, destiny *BankAccount, value float32) {
	origin.CheckIfAccountIsBlocked()
	destiny.CheckIfAccountIsBlocked()

	origin.Withdraw(value)
	destiny.Deposit(value)
}
