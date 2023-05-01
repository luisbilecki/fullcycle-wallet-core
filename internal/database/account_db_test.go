package database

import (
	"database/sql"
	"testing"

	"github.com/luisbilecki/fullcycle-wallet-core/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at date, updated_at date)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance FLOAT, created_at date, updated_at date)")
	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("John Doe", "j@j.com")
}

func (s *AccountDBTestSuite) TearDownTest() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}
