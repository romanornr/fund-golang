package funding

type Fund struct {
	balance int
}

// NewFund function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
	// return a pointer to a new struct without worry
	// if it's on the stack or heap
	return &Fund{
		balance: initialBalance,
	}
}

// Balance Method starts with a receiver, a Fund pointer
func (f *Fund) Balance() int {
	return f.balance
}

// Withdraw an amount from balance
func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
