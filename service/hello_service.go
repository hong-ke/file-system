package service

import (
	"filesystem/entity"
	"filesystem/repository"
	"github.com/pkg/errors"
)

type HelloService struct {
	repositoryFactory repository.RepositoryFactory
}

func NewHelloService(factory repository.RepositoryFactory) (*HelloService, error) {
	if factory == nil {
		return nil, errors.New("nil argument")
	}

	return &HelloService{
		repositoryFactory: factory,
	}, nil
}

func (h *HelloService) Get() (*entity.User, error) {

	tx, err := h.repositoryFactory.NewTransactionScope()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	helloRepository, err := h.repositoryFactory.NewHelloRepository(tx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	get := helloRepository.Get()
	return get, nil

}
