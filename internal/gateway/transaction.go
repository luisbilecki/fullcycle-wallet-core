package gateway

import "github.com/luisbilecki/fullcycle-wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
