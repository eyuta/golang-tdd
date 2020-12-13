package money

// Accessor is a accessor of Money
type Accessor interface {
	Amount() int
	Name() string
}

// Money is a struct that handles money.
type Money struct {
	amount int
	name   string
}

// NewDollar is constructor of Dollar.
func NewDollar(a int) Money {
	return Money{
		amount: a,
		name:   "Dollar",
	}
}

// NewFranc is constructor of Dollar.
func NewFranc(a int) Money {
	return Money{
		amount: a,
		name:   "Franc",
	}
}

// Times multiplies the amount of the receiver by a multiple of the argument
func (m Money) Times(multiplier int) Money {
	return Money{
		amount: m.amount * multiplier,
		name:   m.name,
	}
}

// Equals checks if the amount of the receiver and the argument are the same
func (m Money) Equals(a Accessor) bool {
	return m.amount == a.Amount() && m.name == a.Name()
}

// Amount returns amount field
func (m Money) Amount() int {
	return m.amount
}

// Name returns name field
func (m Money) Name() string {
	return m.name
}
