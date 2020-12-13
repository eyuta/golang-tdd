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

	t.Run("同じ金額のフランが等価である", func(t *testing.T) {
		assert.True(t, money.NewFranc(5).Equals(money.NewFranc(5)))
		assert.False(t, money.NewFranc(5).Equals(money.NewFranc(6)))
	})

	t.Run("同じ金額のドルとフランが等価ではない", func(t *testing.T) {
		assert.False(t, money.NewFranc(5).Equals(money.NewDollar(5)))
	})
	t.Run("通貨テスト", func(t *testing.T) {
		assert.Equal(t, "USD", money.NewDollar(1).Currency())
		assert.Equal(t, "CHF", money.NewFranc(1).Currency())
	})
}
