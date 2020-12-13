package money

// Bank calculates using exchange rates
type Bank struct {
}

// Reduce applies the exchange rate to the argument expression
func (b Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(to)
}
