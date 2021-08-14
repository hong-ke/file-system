package repository

type TransactionScope interface {
	Close() error
	Commit() error
}
