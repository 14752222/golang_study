package global

import (
	"github.com/spf13/viper"
	"demo1/config"
)

type Application struct {
	CofigViper *viper.Viper
	Config config.Configuration
}

var App = new(Application)