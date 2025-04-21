package bank

import (
	"fmt"
	"go-hero/logging"
	"math/rand"
	"sync"
	"time"
)

type BankAccount struct {
	Balance float64
	Owner   string
	mu      sync.RWMutex
	Logger  logging.Logger
}

func (ba *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		ba.mu.Lock()
		defer ba.mu.Unlock()
		ba.Balance += amount
		ba.Logger.Info(fmt.Sprintf("Deposited %.2f. New balance: %.2f", amount, ba.Balance))
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func (ba *BankAccount) Withdraw(amount float64) error {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	if ba.Balance < amount {
		return fmt.Errorf("insufficient balance")
	}

	ba.Balance -= amount
	ba.Logger.Info(fmt.Sprintf("Withdrawn %.2f. New balance: %.2f", amount, ba.Balance))
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return nil
}

func (ba *BankAccount) GetBalance() float64 {
	ba.mu.RLock()
	defer ba.mu.RUnlock()
	fmt.Printf("Available Balance: $%.2f\n", ba.Balance)
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return ba.Balance
}
