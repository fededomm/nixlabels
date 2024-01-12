package main

import (
	"fmt"
	"os"
	"os/exec"
)

func NixOsRebuildCommand(label string) error {
	command := fmt.Sprintf(`sudo NIXOS_LABEL="%s" nixos-rebuild switch`, label)
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v", err)
	}
	return nil
}
