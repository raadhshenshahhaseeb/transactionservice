package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Chain struct {
		PrivateKey string `mapstructure:"PRIVATE_KEY"`
		Endpoint   string `mapstructure:"ENDPOINT"`
	} `mapstructure:"CHAIN"`
	Logger struct {
		Level int    `mapstructure:"LEVEL"`
		Env   string `mapstructure:"ENV"`
	} `mapstructure:"LOGGER"`
}

func Init() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("env.json")
	viper.SetConfigType("json")
	viper.AddConfigPath("../../")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := Config{}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to bootstrap config")
	}

	return &config, nil
}
