package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide an argument")
        os.Exit(1)
    }

    label := os.Args[1]
    err := executeCommand(label)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func executeCommand(label string) error {
	command := fmt.Sprintf(`sudo NIXOS_LABEL="%s" nixos-rebuild switch`, label)
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v", err)
	}
	fmt.Print(string(output))
	return nil
}
