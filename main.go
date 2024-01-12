package main

import (
	"fmt"
	"os"
	"scriptino/commands"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide an argument")
        os.Exit(1)
    }

    label := os.Args[1]
	if err := commands.NixOsRebuildCommand(label); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	
	if err := commands.CommitCommand(label); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := commands.PushCommand(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	

}




