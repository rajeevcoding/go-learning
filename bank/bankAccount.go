package bank

import (
	"fmt"
)

type BankAccount struct {
	Balance float64
	Owner   string
}

func (ba *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		ba.Balance += amount
	}
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if ba.Balance < amount {
		return fmt.Errorf("insufficient balance")
	}

	ba.Balance -= ba.Balance
	return nil
}

func (ba *BankAccount) PrintBalance() {
	fmt.Printf("Available Balance: $%.2f\n", ba.Balance)
}