package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("John Doe Jr", "j@ju.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Jr", client.Name)
	assert.Equal(t, "j@ju.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("", "")
	assert.Error(t, err, "name is required")
}

func TestAddAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
}

func TestAddAccountWithInvalidAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	otherClient, _ := NewClient("Other Jane Doe", "t@t.com")
	account := NewAccount(otherClient)
	err := client.AddAccount(account)
	assert.Error(t, err, "account does not belong to client")
}
