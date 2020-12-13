package money

// Sum 合計は通貨の加算を行います
type Sum struct {
	Augend Money
	Added  Money
}

// Reduce applies the exchange rate to the result of the addition
func (s Sum) Reduce(to string) Money {
	amount := s.Augend.amount + s.Added.amount
	return NewMoney(amount, to)
}
