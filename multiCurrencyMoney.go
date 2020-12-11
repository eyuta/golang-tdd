package main

type Object interface{}

type Dollar struct {
	amount int
}

func (d Dollar) times(multiplier int) Dollar {
	return Dollar{d.amount * multiplier}
}

func (d Dollar) equals(object Object) bool {
	dollar := object.(Dollar)
	return d.amount == dollar.amount
}

type Franc struct {
	amount int
}

func (f Franc) times(multiplier int) Franc {
	return Franc{f.amount * multiplier}
}

func (f Franc) equals(object Object) bool {
	franc := object.(Franc)
	return f.amount == franc.amount
}

func main() {}
