package dig

import (
	"filesystem/clients"
	"filesystem/config"
	"filesystem/repository"
	"filesystem/repository/impl"
	"filesystem/service"
	"go.uber.org/dig"
	"xorm.io/xorm"
)

var container = dig.New()

func initContainer() *dig.Container {
	containerProvide(func() (config.Config, error) {
		return config.Config(*config.GetInstance()), nil
	})
	containerProvide(func() (*xorm.Engine, error) {
		return config.NewDBEngine()
	})
	containerProvide(func(engine *xorm.Engine) (repository.RepositoryFactory, error) {
		return impl.NewRepositoryFactory(engine)
	})
	containerProvide(func(conf config.Config) (*clients.RedisLock, error) {
		return clients.NewRedisLock(conf)
	})
	containerProvide(func(conf config.Config) (*clients.RedisPool, error) {
		return clients.NewRedisClient(conf)
	})
	containerProvide(func(factory repository.RepositoryFactory) (*service.HelloService, error) {
		return service.NewHelloService(factory)
	})
	containerProvide(func(conf config.Config) (*service.UploadService, error) {
		return service.NewUploadService(conf.GetString("upload.path"))
	})
	return container
}
func containerProvide(constructor interface{}, opts ...dig.ProvideOption) {
	err := container.Provide(constructor, opts...)
	if err != nil {
		panic(err)
	}
}

func Container() {
	initContainer()
}
