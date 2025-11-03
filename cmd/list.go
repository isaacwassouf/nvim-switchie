package cmd

import (
	"log"
	"os"

	"github.com/isaacwassouf/nvim-config-switcher/helpers"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		reposPath, err := helpers.GetInstalledCfgsPath()
		if err != nil {
			log.Fatal(err)
		}

		repos, err := os.ReadDir(reposPath)
		if err != nil {
			log.Fatal(err)
		}

		for _, repo := range repos {
			log.Println(repo.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
