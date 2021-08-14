package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	// _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

type Config config

type config struct {
	*viper.Viper
}

var (
	once     sync.Once
	instance *config
	engine   *xorm.Engine
)

func NewDBEngine() (*xorm.Engine, error) {
	return engine, nil
}
func init() {
	once.Do(func() {
		var err error
		var configPath string
		fileName := "config.yml"
		envConfigPath := os.Getenv("CONFIG_PATH")
		if strings.EqualFold(envConfigPath, "") {
			configPath = fmt.Sprintf("./%v", fileName)
			log.Infof("use default path %s", configPath)

		} else {
			log.Infof("find success in env CONFIG_PATH, use %s", envConfigPath)
			configPath = fmt.Sprintf("%v/%v", envConfigPath, fileName)
		}
		instance = &config{viper.New()}
		// 设置默认值
		instance.SetDefault("http.request.timeout.millisecond", 60000)

		// 加载配置文件
		instance.SetConfigType("yaml")
		instance.SetConfigFile(configPath)
		if err = instance.ReadInConfig(); err != nil {
			log.Fatalf("config file read err %s", err.Error())
			panic(err)
		}

		// 根据环境覆盖配置
		environment := os.Getenv("ENVIRONMENT")
		if len(environment) > 0 {
			var path string
			if len(envConfigPath) > 0 {
				path = fmt.Sprintf("%v/%v.yaml", envConfigPath, environment)
			} else {
				path = fmt.Sprintf("./%v.yaml", environment)
			}
			instance.SetConfigFile(path)
			instance.SetConfigType("yaml")
			if err := instance.MergeInConfig(); err != nil {
				log.Fatalf("merge config err:%v", err.Error())
				panic(err)
			}
		}
		// 加载环境变量
		instance.AutomaticEnv()
		instance.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		keys := instance.AllKeys()
		for _, key := range keys {
			fmt.Println(key, ":", instance.Get(key))
		}
		// 初始化xorm 配置
		engine = openDB(
			GetInstance().GetString("base.db.xorm.type"),
			GetInstance().GetString("base.db.xorm.host"),
			GetInstance().GetString("base.db.xorm.username"),
			GetInstance().GetString("base.db.xorm.password"),
			GetInstance().GetString("base.db.xorm.name"),
			GetInstance().GetString("base.db.xorm.port"),
			GetInstance().GetBool("base.db.xorm.showsql"))
	})
}
func GetInstance() *config {
	return instance
}

func openDB(dbType string, host string, username string, password string, name string, port string, showSql bool) *xorm.Engine {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username,
		password,
		host,
		port,
		name)
	engine, err := xorm.NewEngine(dbType, dsn)
	if err != nil {
		log.Errorf("Database connection failed err: %v. Database name: %s", err, name)
		panic(err)
	}

	engine.ShowSQL(showSql)
	return engine
}
