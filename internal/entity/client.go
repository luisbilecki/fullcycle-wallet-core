package entity

import (
	"errors"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string `validate:"required"`
	Email     string `validate:"required"`
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name string, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := client.Validate()
	return client, err
}

func (c *Client) Update(name string, email string) error {
	c.Name = name
	c.Email = email
	return c.Validate()
}

func (c *Client) Validate() error {
	return validator.New().Struct(c)
}

func (c *Client) AddAccount(account *Account) error {
	if account.Client.ID != c.ID {
		return errors.New("account does not belong to client")
	}
	c.Accounts = append(c.Accounts, account)
	return nil
}
