package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiCurrencyMoney(t *testing.T) {
	t.Run("ドルの掛け算が可能である", func(t *testing.T) {
		five := Dollar{5}
		assert.Equal(t, Dollar{10}, five.times(2))
		assert.Equal(t, Dollar{15}, five.times(3))
	})

	t.Run("同じ金額のドルが等価である", func(t *testing.T) {
		assert.True(t, Dollar{5}.equals(Dollar{5}))
		assert.False(t, Dollar{5}.equals(Dollar{6}))
	})

	t.Run("フランの掛け算が可能である", func(t *testing.T) {
		five := Franc{5}
		assert.Equal(t, Franc{10}, five.times(2))
		assert.Equal(t, Franc{15}, five.times(3))
	})

	t.Run("同じ金額のフランが等価である", func(t *testing.T) {
		assert.True(t, Franc{5}.equals(Franc{5}))
		assert.False(t, Franc{5}.equals(Franc{6}))
	})
}
