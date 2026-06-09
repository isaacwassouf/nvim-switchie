package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var repo string
var name string

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(fmt.Errorf("could not determine home directory: %w", err))
		}

		config := filepath.Join(home, ".config", "ghayr", "configs", name)

		// Clone the repo
		cloneCmd := exec.Command("git", "clone", repo, config)
		if err := cloneCmd.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&repo, "repo", "r", "", "Repo that contains the config")
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the config")

	addCmd.MarkFlagRequired("repo")
	addCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(addCmd)
}
