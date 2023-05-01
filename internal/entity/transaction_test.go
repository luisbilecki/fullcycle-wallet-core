package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	clientFrom, _ := NewClient("clientFrom", "c@c.com")
	clientTo, _ := NewClient("clientTo", "d@d.com")

	accountFrom := NewAccount(clientFrom)
	accountTo := NewAccount(clientTo)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	transaction, err := NewTransaction(accountFrom, accountTo, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 1100.0, accountTo.Balance)
	assert.Equal(t, 900.0, accountFrom.Balance)
}

func TestCreateTransactionWithInvalidAccountFrom(t *testing.T) {
	clientFrom, _ := NewClient("clientFrom", "c@c.com")
	clientTo, _ := NewClient("clientTo", "d@d.com")

	accountFrom := NewAccount(clientFrom)
	accountTo := NewAccount(clientTo)

	accountFrom.Credit(90)
	accountTo.Credit(1000)

	_, err := NewTransaction(accountFrom, accountTo, 100)
	assert.Error(t, err, "insufficient funds")
	assert.Equal(t, 90.0, accountFrom.Balance)
	assert.Equal(t, 1000.0, accountTo.Balance)
}
