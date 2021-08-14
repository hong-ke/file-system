package clients

import (
	"filesystem/config"
	"github.com/piaohao/godis"
	"time"
)

type RedisPool struct {
	*godis.Redis
	Prefix string
}

type RedisLock struct {
	*godis.Locker
	Prefix string
}

func NewRedisClient(config config.Config) (*RedisPool, error) {
	pool := initRedisPool(
		config.GetString("base.redis.address"),
		config.GetString("base.redis.password"),
		config.GetInt("base.redis.port"),
		config.GetInt("base.redis.db"),
		config.GetInt("base.redis.maxIdle"),
		config.GetString("base.redis.prefix"),
	)
	return pool, nil
}

func initRedisPool(server, password string, port int, db, maxIdleConn int, prefix string) *RedisPool {
	option := &godis.Option{
		Host:     server,
		Password: password,
		Port:     port,
		Db:       db,
	}
	pool := godis.NewPool(&godis.PoolConfig{MaxIdle: maxIdleConn}, option)
	redis, err := pool.GetResource()
	if err != nil {
		panic(err)
	}
	_, err = redis.Ping()
	if err != nil {
		panic(err)
	}
	redisPool := &RedisPool{
		Redis:  redis,
		Prefix: prefix,
	}
	return redisPool
}

func NewRedisLock(config config.Config) (*RedisLock, error) {
	lock := initRedisLock(
		config.GetString("base.redis.address"),
		config.GetString("base.redis.password"),
		config.GetInt("base.redis.port"),
		config.GetInt("base.redis.db"),
		config.GetInt("base.redis.maxIdle"),
		config.GetString("base.redis.prefix"),
	)
	return lock, nil
}

func initRedisLock(server, password string, port, db, maxIdleConn int, prefix string) *RedisLock {
	option := &godis.Option{
		Host:     server,
		Password: password,
		Port:     port,
		Db:       db,
	}
	lockOption := &godis.LockOption{
		Timeout: 5 * time.Second,
	}
	locker := godis.NewLocker(option, lockOption)
	redisLock := &RedisLock{
		Locker: locker,
		Prefix: prefix,
	}
	return redisLock
}
