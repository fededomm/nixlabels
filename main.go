package main

import (
	"fmt"
	"os"
	"scriptino/commands"
	"scriptino/configurations"
)


func main() {

	conf, err := configurations.ViperConfig()
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
	if err := commands.CdOnNixConfFolger(conf.Path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// copy command inside folder
	if err := commands.BashCopyCommand(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// commit to repository
	if err := commands.CommitCommand(label); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// push to repository
	if err := commands.PushCommand(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}




