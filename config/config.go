package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	ConfigPath = "./script"
)

type Config struct {
	Test          string `json:"test"`
	EsAddress     string `json:"es_address"` //没用
	ThriftAddress string `json:"thrift_address"`
	HasParse      bool   `json:"-"`
}

var config *Config

func GetConfig() *Config {
	return config
}

func InitConfig(configPath string) {
	if config != nil && config.HasParse {
		logrus.Warnf("InitConfig config has parse, config=%+v", config)
		return
	}

	cf := parseConfig(configPath)
	config = cf

	logrus.Infof("InitConfig config=%+v", *cf)
}

func parseConfig(configPath string) *Config {
	cf := Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	if configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		viper.AddConfigPath(ConfigPath)
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("ParseConfig path:%s read error, err=%s", ConfigPath, err.Error()))
	}

	err = viper.Unmarshal(&cf)
	if err != nil {
		panic(fmt.Sprintf("ParseConfig path:%s unmarshl error, err=%s", ConfigPath, err.Error()))
	}

	cf.HasParse = true
	return &cf
}
