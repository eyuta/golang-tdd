package money

// Dollar is a struct that handles dollar money.
type Dollar struct {
	Money
}

// NewDollar is constructor of Dollar.
func NewDollar(a int) Dollar {
	return Dollar{Money{a}}
}

func (d Dollar) times(multiplier int) Dollar {
	return NewDollar(d.amount * multiplier)
}
