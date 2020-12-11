package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMultiCurrencyMoney(t *testing.T) {
	t.Run("$5 * 2 = $10", func(t *testing.T) {
		five := Dollar{5}
		five.times(2)
		assert.Equal(t, 10, five.amount)
	})
}
