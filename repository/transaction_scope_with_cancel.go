package repository

type TransactionScopeWithCancel interface {
	Close() error
	Commit() error
}
