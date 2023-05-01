package client

import (
	"testing"

	mocks "github.com/luisbilecki/fullcycle-wallet-core/internal/gateway/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	clientGateway := &mocks.ClientGatewayMock{}
	clientGateway.On("Save", mock.Anything).Return(nil)

	createClientUseCase := NewCreateClientUseCase(clientGateway)
	input := CreateClientInputDTO{
		Name:  "John Doe",
		Email: "c@c.com",
	}
	output, err := createClientUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Email, output.Email)
	clientGateway.AssertExpectations(t)
	clientGateway.AssertNumberOfCalls(t, "Save", 1)
}
