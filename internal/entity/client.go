package entity

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string `validate:"required"`
	Email     string `validate:"required"`
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
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Validate() error {
	return validator.New().Struct(c)
	// validationErrors := err.(validator.ValidationErrors)
}
