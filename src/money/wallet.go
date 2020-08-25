package money

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	money Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
func (w *Wallet) Deposit(money Bitcoin) {
	w.money += money
}

func (w *Wallet) Balance() Bitcoin {
	return w.money
}

func (w *Wallet) Withdraw(money Bitcoin) error {
	if w.money < money {
		return errors.New("cannot withdraw, insufficient funds")
	}
	w.money -= money
	return nil
}
