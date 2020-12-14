package money

// Expression shows the formula of currency (regardless of the difference in exchange rate)
type Expression interface {
	Reduce(Bank, string) Money
	Plus(Expression) Expression
	Times(int) Expression
}
