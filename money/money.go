package money

import "fmt"

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

func (m Money) String() string {
	return fmt.Sprintf("{Amount: %v, Currency: %v}", m.amount, m.currency)
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
func (m Money) Times(multiplier int) Expression {
	return NewMoney(m.amount*multiplier, m.currency)
}

// Plus adds an argument to the amount of receiver.
func (m Money) Plus(added Expression) Expression {
	return Sum{
		Augend: m,
		Added:  added,
	}
}

// Reduce applies the exchange rate to receiver.
func (m Money) Reduce(b Bank, to string) Money {
	rate := b.Rate(m.currency, to)
	return NewMoney(m.amount/rate, to)
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
