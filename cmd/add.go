package cmd

import (
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

var repo string
var name string

var reposPath string

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		fullRepoPath := path.Join(reposPath, name)

		if _, err := os.Stat(fullRepoPath); !os.IsNotExist(err) {
			log.Fatalf("Repo with name %s already exists", name)
		}

		// Clone the repo
		cloneCmd := exec.Command("git", "clone", repo, fullRepoPath)
		if err := cloneCmd.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	reposPath = path.Join(homeDir, ".config", "switchie", "repos")

	addCmd.Flags().StringVarP(&repo, "repo", "r", "", "Repo that contains the config")
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the config")

	addCmd.MarkFlagRequired("repo")
	addCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(addCmd)
}
