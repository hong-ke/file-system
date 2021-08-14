package dig

import (
	"filesystem/clients"
	"filesystem/service"
	"github.com/pkg/errors"
)

func NewHelloService() (*service.HelloService, error) {
	var hello_Service *service.HelloService
	err := container.Invoke(func(_service *service.HelloService) {
		hello_Service = _service
	})
	return hello_Service, errors.WithStack(err)
}
func NewRedisClient() (*clients.RedisPool, error) {
	var redisPool *clients.RedisPool
	err := container.Invoke(func(_redisPool *clients.RedisPool) {
		redisPool = _redisPool
	})
	return redisPool, errors.WithStack(err)
}

func NewRedisLock() (*clients.RedisLock, error) {
	var redisLock *clients.RedisLock
	err := container.Invoke(func(_redisLock *clients.RedisLock) {
		redisLock = _redisLock
	})
	return redisLock, errors.WithStack(err)
}
