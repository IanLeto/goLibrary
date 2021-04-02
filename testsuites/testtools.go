package testsuites

import (
	"fmt"
	"goLibrary/config"
	"goLibrary/utils"
	"os"
)

var ConfigTestSuit = ConfigSuit{}

type ConfigSuit struct {
	*config.Config
}

func (s ConfigSuit) InitConfigSuit() {
	dir, err := os.Getwd()
	utils.NoErr(err)
	config.InitConfig(fmt.Sprintf(dir+"/config.yaml"))
	//event.Event.Publish("test_config_suit", fmt.Sprintf(dir+"/config.yaml"))

}
