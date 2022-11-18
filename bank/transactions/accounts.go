package accounts

import (
	"fmt"
	"os"
)

type BankAccount struct {
	Owner          string
	Agency         int
	Account        int
	Balance        float32
	AccountBlocked bool
	FirstUse       bool
}

func (Account *BankAccount) CheckIfAccountIsBlocked() {
	if Account.AccountBlocked || Account.FirstUse {
		fmt.Println(Account.Owner + ", you need to unblock your Account to make this transaction")
		os.Exit(-1)
	} else {
		os.Exit(0)
	}
}
func (Account *BankAccount) UnlockAccount() string {
	if Account.AccountBlocked || Account.FirstUse {
		Account.AccountBlocked = false
		Account.FirstUse = false
		return Account.Owner + ", you Account has been unlocked!"
	} else {
		return Account.Owner + ", Account already unlocked"
	}
}

func NewAccount(Owner string, Agency int, Account int, Balance float32) BankAccount {
	return BankAccount{Owner, Agency, Account, Balance, false, true}
}

func (Account *BankAccount) Withdraw(value float32) float32 {
	Account.CheckIfAccountIsBlocked()

	canWithdraw := value <= Account.Balance && value > 0

	if canWithdraw {
		Account.Balance -= value
	} else {
		fmt.Println("Insufficient funds")
	}
	return Account.Balance
}

func (Account *BankAccount) Deposit(value float32) float32 {

	if value > 0 {
		Account.Balance += value
	} else {
		fmt.Println("Invalid value")
	}
	return Account.Balance
}

func Transfer(origin *BankAccount, destiny *BankAccount, value float32) {
	origin.CheckIfAccountIsBlocked()
	destiny.CheckIfAccountIsBlocked()

	origin.Withdraw(value)
	destiny.Deposit(value)
}
