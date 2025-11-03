package cmd

import (
	"log"
	"os/exec"
	"time"

	getter "github.com/hashicorp/go-getter"
	"github.com/spf13/cobra"
)

var gitGetter = &getter.GitGetter{
	Timeout: 5 * time.Minute,
}
var repo string

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		cmdClone := exec.Command("git", "clone", repo, "./somethinging")
		if err := cmdClone.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&repo, "repo", "r", "", "Repo that contains the config")
	addCmd.MarkFlagRequired("repo")

	rootCmd.AddCommand(addCmd)
}
