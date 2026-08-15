package pointeranderrors

import (
	"errors"
	"fmt"
)

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type wallet struct {
	balance Bitcoin
}

func (w *wallet) Balance() Bitcoin {
	return w.balance
}

func (w *wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}
