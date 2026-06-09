package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:  "use",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configName := args[0]

		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("could not determine home directory: %v", err)
		}

		configDir := filepath.Join(home, ".config", "ghayr", "configs", configName)
		if _, err := os.Stat(configDir); err != nil {
			if os.IsNotExist(err) {
				log.Fatalf("config '%s' does not exist", configName)
			}
			log.Fatalf("could not check config directory: %v", err)
		}

		currentFile := filepath.Join(home, ".config", "ghayr", ".current")
		if err := os.WriteFile(currentFile, []byte(configName), 0644); err != nil {
			log.Fatalf("could not write to .current file: %v", err)
		}

		fmt.Printf("Switched to config '%s'\n", configName)

	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
