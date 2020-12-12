package money


// Franc is a struct that handles franc money.
type Franc struct {
	Money
}

// NewFranc is constructor of Dollar.
func NewFranc(a int) Franc {
	return Franc{Money{a}}
}

func (f Franc) times(multiplier int) Franc {
	fmt.Println(multiplier)
	return NewFranc(f.amount * multiplier)
}
