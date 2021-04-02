package config

import (
	"fmt"
	"os"
)

func InitConfig(path string) {
	var (
		err error
	)
	if path == "" {
		dir, _ := os.Getwd()
		path = fmt.Sprintf(dir + "config.yaml")
	}
	Viper.SetConfigFile(path)
	err = Viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	NewBaseConfig()
}

func NewBaseConfig() {
	BaseConfig = &Config{
		Name:        Viper.GetString("name"),
		Port:        Viper.GetString("addr"),
		RunMode:     Viper.GetString("run_mode"),
		MySqlConfig: NewMySqlConfig(),
		RedisConfig: NewRedisConfig(),
	}
}

func NewMySqlConfig() MySqlConfig {
	return MySqlConfig{
		Address:  Viper.GetString("MySql.address"),
		Port:     Viper.GetString("MySql.port"),
		User:     Viper.GetString("MySql.user"),
		Password: Viper.GetString("MySql.password"),
	}
}

func NewRedisConfig() RedisConfig {
	return RedisConfig{
		Address:  Viper.GetString("redis.address"),
		Port:     Viper.GetString("redis.port"),
		IsMaster: Viper.GetBool("redis.is_master"),
	}
}
