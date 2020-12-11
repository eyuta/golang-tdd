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
func main() {}
