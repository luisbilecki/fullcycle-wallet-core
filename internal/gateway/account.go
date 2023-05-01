package gateway

import "github.com/luisbilecki/fullcycle-wallet-core/internal/entity"

type AccountGateway interface {
	Save(client *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
