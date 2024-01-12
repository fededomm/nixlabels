package main

import (
	"fmt"
	"os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide an argument")
        os.Exit(1)
    }

    label := os.Args[1]
	err := NixOsRebuildCommand(label)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}




