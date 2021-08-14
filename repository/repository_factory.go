package repository

import "context"

type RepositoryFactory interface {
	NewTransactionScope() (TransactionScope, error)
	NewTransactionScopeWithCancel() (TransactionScopeWithCancel, context.CancelFunc, error)
	NewHelloRepository(tx TransactionScope) (HelloRepository, error)
}
