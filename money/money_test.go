package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiCurrencyMoney(t *testing.T) {
	t.Run("ドルの掛け算が可能である", func(t *testing.T) {
		five := NewDollar(5)
		assert.Equal(t, NewDollar(10), five.times(2))
		assert.Equal(t, NewDollar(15), five.times(3))
	})

	t.Run("同じ金額のドルが等価である", func(t *testing.T) {
		assert.True(t, NewDollar(5).equals(NewDollar(5)))
		assert.False(t, NewDollar(5).equals(NewDollar(6)))
	})

	t.Run("フランの掛け算が可能である", func(t *testing.T) {
		five := NewFranc(5)
		assert.Equal(t, NewFranc(10), five.times(2))
		assert.Equal(t, NewFranc(15), five.times(3))
	})

	t.Run("同じ金額のフランが等価である", func(t *testing.T) {
		assert.True(t, NewFranc(5).equals(NewFranc(5)))
		assert.False(t, NewFranc(5).equals(NewFranc(6)))
	})
}
