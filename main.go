package main

import (
	"fmt"
	"os"
	"scriptino/commands"

	"github.com/spf13/viper"
)

type Config struct {
	Path string `mapstructure:"path" yaml:"path"`
}

var config *Config

func main() {

	// Read Config File and unmarshal to struct
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/Desktop/nixlabels")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	if len(os.Args) < 2 {
        fmt.Println("Please provide an argument")
        os.Exit(1)
    }

	// NixOsRebuildCommand
    label := os.Args[1]
	if err := commands.NixOsRebuildCommand(label); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

	// cd on destination folder
	if err := commands.CdOnNixConfFolger(config.Path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// copy command inside folder
	if err := commands.BashCopyCommand(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 
	if err := commands.CommitCommand(label); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//
	if err := commands.PushCommand(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}




