package money

// Bank calculates using exchange rates
type Bank struct {
	rates map[Pair]int
}

// Pair associates two currencies
type Pair struct {
	from, to string
}

// NewBank is a constructor of Bank
func NewBank() Bank {
	b := Bank{}
	b.rates = make(map[Pair]int)
	return b
}

// Reduce applies the exchange rate to the argument expression
func (b *Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(*b, to)
}

// AddRate adds exchange rate
func (b *Bank) AddRate(from, to string, rate int) {
	b.rates[Pair{from: from, to: to}] = rate
}

// Rate adds exchange rate
func (b *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	p := Pair{from: from, to: to}
	return b.rates[p]
}
