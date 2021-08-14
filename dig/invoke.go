package dig

import (
	"filesystem/clients"
	"filesystem/service"
	"github.com/pkg/errors"
)

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

func NewUploadService() (*service.UploadService, error) {
	var upload_Service *service.UploadService
	err := container.Invoke(func(_service *service.UploadService) {
		upload_Service = _service
	})
	return upload_Service, errors.WithStack(err)
}
