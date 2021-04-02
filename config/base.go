package config

import "github.com/spf13/viper"

var Viper = viper.New()

type Config struct {
	Name         string `yaml:"name"`
	Address      string `yaml:"address"`
	Port         string `yaml:"port"`
	RunMode      string `yaml:"run_mode"`
	LoggerConfig LoggerConfig
	MySqlConfig  MySqlConfig
	RedisConfig  RedisConfig
}

type LoggerConfig struct {
	LogPath       string `json:"log_path"`
	RotationTime  int    `json:"rotation_time"`
	RotationCount int    `json:"rotation_count"`
}

type RedisConfig struct {
	Address  string `json:"address"`
	Port     string `json:"port"`
	IsMaster bool   `json:"is_master"`
}

type MySqlConfig struct {
	Address  string `json:"address"`
	Port     string `json:"port"`
	Password string `json:"password"`
	User     string `json:"user"`
}

func SubInitConfigEvent() {

}
