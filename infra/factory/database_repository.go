package factory

import (
	"database/sql"

	"github.com/brenoxavier48/imersaofc-gateway/domain/repository"
	repo "github.com/brenoxavier48/imersaofc-gateway/infra/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repo.NewTransactionRepositoryDB(r.DB)
}
