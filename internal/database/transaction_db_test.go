package database

import (
	"database/sql"
	"testing"

	"github.com/luisbilecki/fullcycle-wallet-core/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	clientFrom    *entity.Client
	clientTo      *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at date, updated_at date)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance FLOAT, created_at date, updated_at date)")
	db.Exec("CREATE TABLE transactions (id VARCHAR(255), account_id_from VARCHAR(255), account_id_to VARCHAR(255), amount FLOAT, created_at date, updated_at date)")
	s.clientFrom, _ = entity.NewClient("John Doe", "j@j.com")
	s.clientTo, _ = entity.NewClient("John Doe", "b@b.com")

	accountFrom := entity.NewAccount(s.clientFrom)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(s.clientTo)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownTest() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE transactions")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
