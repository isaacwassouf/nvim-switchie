package cmd

import (
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/isaacwassouf/nvim-config-switcher/helpers"
	"github.com/spf13/cobra"
)

var repo string
var name string

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		reposPath, err := helpers.GetInstalledCfgsPath()
		if err != nil {
			log.Fatal(err)
		}

		newCfgPath := path.Join(reposPath, name)
		if _, err := os.Stat(newCfgPath); !os.IsNotExist(err) {
			log.Fatalf("Repo with name %s already exists", name)
		}

		// Clone the repo
		cloneCmd := exec.Command("git", "clone", repo, newCfgPath)
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
