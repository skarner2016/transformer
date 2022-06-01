package config

import (
	"github.com/spf13/viper"
	"path"
	"runtime"
)

var VipConfig *viper.Viper

func InitConfig() () {
	if VipConfig != nil{
		return
	}

	viperConfig := viper.New()

	// 配置文件名称
	viperConfig.SetConfigName("app")

	// 配置文件路径
	configPath := getConfigPath()
	viperConfig.AddConfigPath(configPath)

	// 配置文件类型
	viperConfig.SetConfigType("toml")

	if err := viperConfig.ReadInConfig(); err != nil {
		panic(err)
	}

	VipConfig = viperConfig
}

func getConfigPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	currentPath := path.Dir(filename)

	return currentPath + "/../../conf"
}
