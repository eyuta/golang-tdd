package main

type AmountGetter interface {
	getAmount() int
}

type Money struct {
	amount int
}

func (m Money) equals(a AmountGetter) bool {
	return m.getAmount() == a.getAmount()
}

func (m Money) getAmount() int {
	return m.amount
}

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

type Franc struct {
	Money
}

// NewFranc is constructor of Dollar.
func NewFranc(a int) Franc {
	return Franc{Money{a}}
}

func (f Franc) times(multiplier int) Franc {
	return NewFranc(f.amount * multiplier)
}

func main() {}
