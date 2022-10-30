package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	AppMode  string
	HttpPort string
}

type Config struct {
	Server ServerConfig
}

var Conf Config

func init() {
	viper.SetConfigName("service")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	// 环境变量
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Conf)
}
