package commands

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

func CommitCommand(label string) error {
	command := fmt.Sprintf(`git add . && git commit -m "Generation label: %s"`, label)
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v", err)
	}
	return nil
}

func PushCommand() error {
	cmd := exec.Command("bash", "-c", "git push")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v", err)
	}
	return nil
}
