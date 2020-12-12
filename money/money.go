package money

// AmountGetter is a wrapper of amount.
type AmountGetter interface {
	getAmount() int
}

// Money is a struct that handles money.
type Money struct {
	amount int
}

func (m Money) equals(a AmountGetter) bool {
	return m.getAmount() == a.getAmount()
}

func (m Money) getAmount() int {
	return m.amount
}
