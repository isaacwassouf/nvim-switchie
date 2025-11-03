package cmd

import (
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

var repo string
var toolReposPath string

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := path.Join(toolReposPath, "testing")
		cmdClone := exec.Command("git", "clone", repo, repoPath)
		if err := cmdClone.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	toolReposPath = path.Join(homeDir, ".config", "switchie", "repos")

	addCmd.Flags().StringVarP(&repo, "repo", "r", "", "Repo that contains the config")
	addCmd.MarkFlagRequired("repo")

	rootCmd.AddCommand(addCmd)
}
