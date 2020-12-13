package money

// Dollar is a struct that handles dollar money.
type Dollar struct {
	Money
}

// NewDollar is constructor of Dollar.
func NewDollar(a int) Dollar {
	return Dollar{Money{amount: a, name: "Dollar"}}
}

// Times multiplies the amount of the receiver by a multiple of the argument
func (d Dollar) Times(multiplier int) Dollar {
	return NewDollar(d.amount * multiplier)
}
