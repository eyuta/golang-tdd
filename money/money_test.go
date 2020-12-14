package money_test

import (
	"testing"

	"github.com/eyuta/golang-tdd/money"
	"github.com/stretchr/testify/assert"
)

func TestMultiCurrencyMoney(t *testing.T) {
	t.Run("ドルの掛け算が可能である", func(t *testing.T) {
		five := money.NewDollar(5)
		assert.Equal(t, money.NewDollar(10), five.Times(2))
		assert.Equal(t, money.NewDollar(15), five.Times(3))
	})

	t.Run("同じ金額のドルが等価である", func(t *testing.T) {
		assert.True(t, money.NewDollar(5).Equals(money.NewDollar(5)))
		assert.False(t, money.NewDollar(5).Equals(money.NewDollar(6)))
	})

	t.Run("フランの掛け算が可能である", func(t *testing.T) {
		five := money.NewFranc(5)
		assert.Equal(t, money.NewFranc(10), five.Times(2))
		assert.Equal(t, money.NewFranc(15), five.Times(3))
	})

	t.Run("同じ金額のドルとフランが等価ではない", func(t *testing.T) {
		assert.False(t, money.NewFranc(5).Equals(money.NewDollar(5)))
	})
	t.Run("通貨テスト", func(t *testing.T) {
		assert.Equal(t, "USD", money.NewDollar(1).Currency())
		assert.Equal(t, "CHF", money.NewFranc(1).Currency())
	})
	t.Run("ドル同士の足し算が可能である", func(t *testing.T) {
		five := money.NewDollar(5)
		sum := five.Plus(five)
		bank := money.NewBank()
		reduced := bank.Reduce(sum, "USD")
		assert.Equal(t, money.NewDollar(10), reduced)
	})
	t.Run("ドル同士の足し算が可能である", func(t *testing.T) {
		five := money.NewDollar(5)
		result := five.Plus(five)
		sum := result.(money.Sum)
		assert.Equal(t, five, sum.Augend)
		assert.Equal(t, five, sum.Added)
	})
	t.Run("Sumで足されるお金の通貨が同じなら、足し算の結果が同じになる", func(t *testing.T) {
		sum := money.Sum{
			Augend: money.NewDollar(3),
			Added:  money.NewDollar(4),
		}
		bank := money.NewBank()
		result := bank.Reduce(sum, "USD")
		assert.Equal(t, money.NewDollar(7), result)
	})
	t.Run("moneyをreduceしても、reduceに渡す通貨が同じであれば同じ値が返る", func(t *testing.T) {
		bank := money.NewBank()
		result := bank.Reduce(money.NewDollar(1), "USD")
		assert.Equal(t, money.NewDollar(1), result)
	})
	t.Run("1 CHF = $2", func(t *testing.T) {
		bank := money.NewBank()
		bank.AddRate("CHF", "USD", 2)
		result := bank.Reduce(money.NewFranc(2), "USD")
		assert.Equal(t, money.NewDollar(1), result)
	})
	t.Run("通貨が同じ場合はレートが1:1になる", func(t *testing.T) {
		bank := money.NewBank()
		assert.Equal(t, 1, bank.Rate("USD", "USD"))
	})
	t.Run("$5 + 10 CHF = $10 (レートが2:1の場合)", func(t *testing.T) {
		fiveBucks := money.Expression(money.NewDollar(5))
		tenFrancs := money.Expression(money.NewFranc(10))
		bank := money.NewBank()
		bank.AddRate("CHF", "USD", 2)
		result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
		assert.Equal(t, money.NewDollar(10), result)
	})
	t.Run("$5 + 10 CHF + $5 = $15 をSum structを使って行う", func(t *testing.T) {
		fiveBucks := money.Expression(money.NewDollar(5))
		tenFrancs := money.Expression(money.NewFranc(10))
		bank := money.NewBank()
		bank.AddRate("CHF", "USD", 2)
		sum := money.Sum{Augend: fiveBucks, Added: tenFrancs}.Plus(fiveBucks)
		result := bank.Reduce(sum, "USD")
		assert.Equal(t, money.NewDollar(15), result)
	})
	t.Run("($5 + 10 CHF) * 2 = $20 をSum structを使って行う", func(t *testing.T) {
		fiveBucks := money.Expression(money.NewDollar(5))
		tenFrancs := money.Expression(money.NewFranc(10))
		bank := money.NewBank()
		bank.AddRate("CHF", "USD", 2)
		sum := money.Sum{Augend: fiveBucks, Added: tenFrancs}.Times(2)
		result := bank.Reduce(sum, "USD")
		assert.Equal(t, money.NewDollar(20), result)
	})
}
