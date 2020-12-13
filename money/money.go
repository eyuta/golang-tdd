package money

// Accessor is a accessor of Money
type Accessor interface {
	Amount() int
	Currency() string
}

// Money is a struct that handles money.
type Money struct {
	amount   int
	currency string
}

// NewMoney is constructor of Money.
func NewMoney(a int, c string) Money {
	return Money{
		amount:   a,
		currency: c,
	}
}

// NewDollar is constructor of Dollar.
func NewDollar(a int) Money {
	return NewMoney(a, "USD")
}

// NewFranc is constructor of Dollar.
func NewFranc(a int) Money {
	return NewMoney(a, "CHF")
}

// Times multiplies the amount of the receiver by a multiple of the argument
func (m Money) Times(multiplier int) Money {
	return Money{
		amount:   m.amount * multiplier,
		currency: m.currency,
	}
}

// Equals checks if the amount of the receiver and the argument are the same
func (m Money) Equals(a Accessor) bool {
	return m.amount == a.Amount() && m.currency == a.Currency()
}

// Amount returns amount field
func (m Money) Amount() int {
	return m.amount
}

// Currency returns name field
func (m Money) Currency() string {
	return m.currency
}
