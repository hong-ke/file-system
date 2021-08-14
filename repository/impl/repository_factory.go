package impl

import (
	"context"
	"filesystem/repository"
	"github.com/pkg/errors"
	"xorm.io/xorm"
)

type RepositoryFactory struct {
	engine *xorm.Engine
}

func NewRepositoryFactory(engine *xorm.Engine) (*RepositoryFactory, error) {
	if engine == nil {
		return nil, errors.New("Nil arguments")
	}

	return &RepositoryFactory{
		engine: engine,
	}, nil
}

func (rf *RepositoryFactory) NewTransactionScope() (repository.TransactionScope, error) {
	session := rf.engine.NewSession()
	err := session.Begin()
	return session, errors.WithStack(err)
}

func (rf *RepositoryFactory) NewTransactionScopeWithCancel() (repository.TransactionScopeWithCancel, context.CancelFunc, error) {
	session := rf.engine.NewSession()
	ctx, cancel := context.WithCancel(context.Background())
	session.Context(ctx)

	err := session.Begin()
	return session, cancel, errors.WithStack(err)
}

func (rf *RepositoryFactory) NewHelloRepository(tx repository.TransactionScope) (repository.HelloRepository, error) {
	if session, ok := tx.(*xorm.Session); ok {
		return &HelloRepository{
			Session: session,
		}, nil
	}
	return nil, errors.New("invalid arguments, *xorm.Session needed")
}
