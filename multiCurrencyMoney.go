package main

type Dollar struct {
	amount int
}

func (d *Dollar) times(multiplier int) {
	d.amount *= multiplier
}

func main() {}
