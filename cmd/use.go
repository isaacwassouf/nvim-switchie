package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use: "use",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("using config")
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
