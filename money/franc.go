package money

// Franc is a struct that handles franc money.
type Franc struct {
	Money
}

// NewFranc is constructor of Dollar.
func NewFranc(a int) Franc {
	return Franc{Money{amount: a, name: "Franc"}}
}

// Times multiplies the amount of the receiver by a multiple of the argument
func (f Franc) Times(multiplier int) Franc {
	return NewFranc(f.amount * multiplier)
}
