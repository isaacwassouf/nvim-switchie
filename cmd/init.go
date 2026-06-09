package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/isaacwassouf/ghayr/internal/scripts"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		if err := bootstrapConfig(); err != nil {
			panic(err)
		}

		if err := addShellFunc(); err != nil {
			panic(err)
		}
	},
}

func bootstrapConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not determine home directory: %w", err)
	}

	configDir := filepath.Join(home, ".config", "ghayr", "configs")

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("could not create config directory: %w", err)
	}

	currentFile := filepath.Join(home, ".config", "ghayr", ".current")
	if _, err := os.Stat(currentFile); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("could not check .current file: %w", err)
		}
		f, err := os.Create(currentFile)
		if err != nil {
			return fmt.Errorf("could not create .current file: %w", err)
		}
		defer f.Close()
	}

	return nil
}

func addShellFunc() error {
	rcFilePath, err := getRCShellPath()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(rcFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprint(f, scripts.Nvims)
	return err
}

func getRCShellPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine home directory: %w", err)
	}

	shell := os.Getenv("SHELL")
	switch {
	case strings.HasSuffix(shell, "zsh"):
		return filepath.Join(home, ".zshrc"), nil
	case strings.HasSuffix(shell, "bash"):
		return filepath.Join(home, ".bashrc"), nil
	default:
		return "", fmt.Errorf("unsupported shell: %s", shell)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
