package main

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	GitHubUrl string `mapstructure:"github-url" yaml:"github-url"`
}

var config *AppConfig

func ReadConfig() (*AppConfig, error) {
	viper.SetConfigName("config") 
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") 
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err 
	}
	return config, nil
}