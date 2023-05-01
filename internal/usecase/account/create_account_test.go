package account

import (
	"testing"

	"github.com/luisbilecki/fullcycle-wallet-core/internal/entity"
	mocks "github.com/luisbilecki/fullcycle-wallet-core/internal/gateway/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "j@l.com")
	clientGateway := &mocks.ClientGatewayMock{}
	clientGateway.On("Get", mock.Anything).Return(client, nil)

	accountGateway := &mocks.AccountGatewayMock{}
	accountGateway.On("Save", mock.Anything).Return(nil)

	createAccountUseCase := NewCreateAccountUseCase(accountGateway, clientGateway)
	input := CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := createAccountUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	clientGateway.AssertExpectations(t)
	clientGateway.AssertNumberOfCalls(t, "Get", 1)
	accountGateway.AssertExpectations(t)
	accountGateway.AssertNumberOfCalls(t, "Save", 1)
}
