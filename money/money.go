package money

// AmountGetter is a wrapper of amount.
type AmountGetter interface {
	getAmount() int
}

// Money is a struct that handles money.
type Money struct {
	amount int
}

// Equals checks if the amount of the receiver and the argument are the same
func (m Money) Equals(a AmountGetter) bool {
	return m.getAmount() == a.getAmount()
}

func (m Money) getAmount() int {
	return m.amount
}
