package money

// Sum adds currency
type Sum struct {
	Augend Expression
	Added  Expression
}

// Reduce applies the exchange rate to the result of the addition
func (s Sum) Reduce(b Bank, to string) Money {
	amount := s.Augend.Reduce(b, to).amount + s.Added.Reduce(b, to).amount
	return NewMoney(amount, to)
}

// Plus adds an argument to the amount of receiver.
func (s Sum) Plus(added Expression) Expression {
	return Sum{
		Augend: s,
		Added:  added,
	}
}

// Times multiplies the amount of the receiver by a multiple of the argument
func (s Sum) Times(multiplier int) Expression {
	return Sum{
		Augend: s.Augend.Times(multiplier),
		Added:  s.Added.Times(multiplier),
	}
}
