package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	AddItem("user-a", "101", 1)
	cart := GetCart("user-a")
	assert.Equal(t, cart.UserId, "user-a")
	assert.Equal(t, len(cart.Items), 1)
	AddItem("user-a", "101", 1)
	assert.Equal(t, cart.Items[0].Quantity, int32(2))
	assert.Equal(t, cart.Items[0].ProductId, "101")
}
