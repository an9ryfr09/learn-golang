package util

import (
	"errors"
	"fmt"
)

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

//Bitcoin ...
type Bitcoin int

//Wallet ...
type Wallet struct {
	balance Bitcoin
}

//Stringer ...
type Stringer interface {
	String() string
}

//Deposit ...
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Withdraw ...
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount <= w.balance {
		w.balance -= amount
		return nil
	}
	return InsufficientFundsError
}

//Balance ..
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
