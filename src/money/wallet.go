package money

type Wallet struct {
	money int
}

func (w *Wallet) Deposit(money int) {
	w.money += money
}

func (w *Wallet) Balance() int {
	return w.money
}
