package configurations

import (
	"fmt"

	"github.com/spf13/viper"
)


type Config struct {
	Path string `mapstructure:"path" yaml:"path"`
}

var config *Config

func ViperConfig() (*Config, error) {
	// Read Config File and unmarshal to struct
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/Desktop/nixlabels")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return config, nil
}