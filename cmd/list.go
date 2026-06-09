package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("could not determine home directory: %v", err)
		}

		configsDir := filepath.Join(home, ".config", "ghayr", "configs")
		entries, err := os.ReadDir(configsDir)
		if err != nil {
			log.Fatalf("could not read configs directory: %v", err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				fmt.Println(entry.Name())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
