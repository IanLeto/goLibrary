package viperDemo

import "github.com/spf13/viper"

var Conf = viper.New()

type Configuration struct {
	Backend *Backend
}
type Backend struct {
	MySql *MySqlConf
	Redis *RedisConf
}

type MySqlConf struct {
	Address string
	Port    string
}
type RedisConf struct {
	Address string
	Port    string
}

func InitConf() {
	//viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	//viper.SetConfigName("config") // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	//viper.AddConfigPath(".")               // 还可以在工作目录中查找配置

	Conf.SetConfigFile("./config.yaml")
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项

	if err := Conf.ReadInConfig(); err != nil {
		panic(err)
	}
}

func InitConfiguration() *Configuration {
	return &Configuration{
		Backend: NewBackendConfig(),
	}
}

func NewBackendConfig() *Backend {
	return &Backend{
		MySql: NewMysqlConf(),
	}
}

func NewMysqlConf() *MySqlConf {
	return &MySqlConf{
		Address: Conf.GetString("backend.mysql.address"),
		Port:    Conf.GetString("backend.mysql.port"),
	}
}

func init() {
	InitConf()
}
