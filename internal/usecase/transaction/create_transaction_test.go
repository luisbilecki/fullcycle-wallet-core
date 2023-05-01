package transaction

import (
	"testing"

	"github.com/luisbilecki/fullcycle-wallet-core/internal/entity"
	mocks "github.com/luisbilecki/fullcycle-wallet-core/internal/gateway/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	clientFrom, _ := entity.NewClient("John Doe", "j@j.com")
	accountFrom := entity.NewAccount(clientFrom)
	accountFrom.Credit(100)

	clientTo, _ := entity.NewClient("John Doe To", "l@l.com")
	accountTo := entity.NewAccount(clientTo)
	accountTo.Credit(100)

	accountGateway := &mocks.AccountGatewayMock{}
	accountGateway.On("FindByID", accountFrom.ID).Return(accountFrom, nil)
	accountGateway.On("FindByID", accountTo.ID).Return(accountTo, nil)

	transactionGateway := &mocks.TransactionGatewayMock{}
	transactionGateway.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIDFrom: accountFrom.ID,
		AccountIDTo:   accountTo.ID,
		Amount:        10,
	}

	createTransactionUseCase := NewCreateTransactionUseCase(transactionGateway, accountGateway)
	output, err := createTransactionUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	accountGateway.AssertExpectations(t)
	accountGateway.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGateway.AssertExpectations(t)
	transactionGateway.AssertNumberOfCalls(t, "Create", 1)
}
