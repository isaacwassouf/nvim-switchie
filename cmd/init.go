package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Could not get the user home directory")
		}

		toolConfigDir := path.Join(homeDir, ".config", "switchie")
		// check if the config directory does not exists, and create it
		if _, err = os.Stat(toolConfigDir); os.IsNotExist(err) {
			err = os.MkdirAll(toolConfigDir, 0755)
			if err != nil {
				log.Fatalf("Could not create the directory: %s", toolConfigDir)
			}

			fmt.Printf("Successfully created: %s \n", toolConfigDir)
		}

		toolRepoDir := path.Join(homeDir, ".config", "switchie", "repos")
		// create repos dir
		if err = os.Mkdir(toolRepoDir, 0755); err != nil {
			log.Fatalf("Could not create repos directory: %s", toolRepoDir)
		}

		fmt.Printf("Successfully created: %s \n", toolRepoDir)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
