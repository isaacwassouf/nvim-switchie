package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "ghayr",
	Run: func(cmd *cobra.Command, args []string) {
		selfCmd, err := os.Executable()
		if err != nil {
			log.Fatalf("could not determine executable path: %v", err)
		}

		var out bytes.Buffer
		fzfCmd := exec.Command("fzf", "--border", "--height=50%", "--prompt=Select Neovim config: ")
		fzfCmd.Stdin = os.Stdin
		fzfCmd.Stdout = &out
		fzfCmd.Stderr = os.Stderr
		fzfCmd.Env = append(os.Environ(),
			fmt.Sprintf("FZF_DEFAULT_COMMAND=%s list", selfCmd),
		)

		if err := fzfCmd.Run(); err != nil {
			var exitErr *exec.ExitError
			if !errors.As(err, &exitErr) {
				log.Fatalf("could not run fzf: %v", err)
			}
		}

		choice := strings.TrimSpace(out.String())
		if choice == "" {
			log.Fatalf("no selection made")
		}

		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("could not determine home directory: %v", err)
		}

		choiceConfigDir := filepath.Join(home, ".config", "ghayr", "configs", choice)
		if _, err := os.Stat(choiceConfigDir); err != nil {
			if os.IsNotExist(err) {
				log.Fatalf("config '%s' does not exist", choice)
			}
			log.Fatalf("could not check config directory: %v", err)
		}

		currentFile := filepath.Join(home, ".config", "ghayr", ".current")
		if err := os.WriteFile(currentFile, []byte(choice), 0644); err != nil {
			log.Fatalf("could not write to .current file: %v", err)
		}

		fmt.Printf("Switched to config '%s'\n", choice)

	},
}

func Execute() error {
	return rootCmd.Execute()
}
