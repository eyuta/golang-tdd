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

// Equals checks if the amount of the receiver and the argument are the same
func (m Money) Equals(a Accessor) bool {
	return m.Amount() == a.Amount() && m.Name() == a.Name()
}

// Amount returns amount field
func (m Money) Amount() int {
	return m.amount
}

// Name returns name field
func (m Money) Name() string {
	return m.name
}
