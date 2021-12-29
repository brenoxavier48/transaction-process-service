package factory

import "github.com/brenoxavier48/imersaofc-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
