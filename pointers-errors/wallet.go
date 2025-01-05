package wallet

import (
	"errors"
	"fmt"
)

var ErrorInsufficientFunds = errors.New("insufficient funds")

type Wallet struct {
	balance Bitcoin
}

type Bitcoin float64

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance() {
		return ErrorInsufficientFunds
	}
	w.balance -= amount
	return nil
}
