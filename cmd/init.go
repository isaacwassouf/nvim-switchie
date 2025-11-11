package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/isaacwassouf/nvim-config-switcher/configs"
	"github.com/isaacwassouf/nvim-config-switcher/helpers"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {

		toolConfigDir, err := helpers.PathFromUserCfg(configs.ToolCfgDir)
		if err != nil {
			log.Fatalf("Could not get the tool config directory: %v", err)
		}
		// check if the config directory does not exists, and create it
		if _, err = os.Stat(toolConfigDir); os.IsNotExist(err) {
			err = os.MkdirAll(toolConfigDir, 0755)
			if err != nil {
				log.Fatalf("Could not create the directory: %s", toolConfigDir)
			}

			fmt.Printf("Successfully created: %s \n", toolConfigDir)
		}

		toolRepoDir, err := helpers.PathFromUserCfg(configs.ToolCfgDir, configs.AddCfgsDir)
		if err != nil {
			log.Fatalf("Could not get the tool repos directory: %v", err)
		}
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
