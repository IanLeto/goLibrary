package Config

import "github.com/spf13/viper"

var VIPER = viper.New()

type Config struct {
	MySql MySqlConfig
}

type MySqlConfig struct {
}

func InitConfig() {
	VIPER = viper.New()
	VIPER.SetConfigType("yaml")
	VIPER.AddConfigPath("./../Config/conf.yaml")
	err := VIPER.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
