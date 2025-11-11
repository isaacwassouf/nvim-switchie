package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/isaacwassouf/nvim-config-switcher/configs"
	"github.com/isaacwassouf/nvim-config-switcher/helpers"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {

		if hasBeenInitialized() {
			log.Fatal("The tool has already been initialized.")
		}

		toolCfgPath, err := helpers.PathFromUserCfg(configs.ToolCfgDir)
		if err != nil {
			log.Fatalf("Could not get the tool config directory: %v", err)
		}

		// check if the config directory does not exists, and create it
		if _, err = os.Stat(toolCfgPath); os.IsNotExist(err) {
			err = os.MkdirAll(toolCfgPath, 0755)
			if err != nil {
				log.Fatalf("Could not create the directory: %s", toolCfgPath)
			}

			fmt.Printf("Successfully created: %s \n", toolCfgPath)
		}

		addCfgsPath, err := helpers.PathFromUserCfg(configs.ToolCfgDir, configs.AddCfgsDir)
		if err != nil {
			log.Fatalf("Could not get the tool repos directory: %v", err)
		}

		// create repos dir
		if err = os.Mkdir(addCfgsPath, 0755); err != nil {
			log.Fatalf("Could not create repos directory: %s", addCfgsPath)
		}

		fmt.Printf("Successfully created: %s \n", addCfgsPath)

		// check if there is a nvim config.
		// If there is, move it to the repos dir, and create a symlink

		nvimCfgPath, err := helpers.PathFromUserCfg("nvim")
		if err != nil {
			log.Fatalf("Could not get nvim config directory: %v", err)
		}

		if _, err = os.Stat(nvimCfgPath); err == nil {
			destPath := path.Join(addCfgsPath, "default-nvim-config")
			err = os.Rename(nvimCfgPath, destPath)
			if err != nil {
				log.Fatalf("Could not move existing nvim config to repos dir: %v", err)
			}

			err = os.Symlink(destPath, nvimCfgPath)
			if err != nil {
				log.Fatalf("Could not create symlink for nvim config: %v", err)
			}
		} else {
			fmt.Printf("error checking nvim config path: %v\n", err)
		}
	},
}

func hasBeenInitialized() bool {
	toolCfgPath, err := helpers.PathFromUserCfg(configs.ToolCfgDir)
	if err != nil {
		log.Fatalf("Could not get the tool config directory: %v", err)
	}

	if _, err = os.Stat(toolCfgPath); os.IsNotExist(err) {
		return false
	}

	return true
}

func init() {
	rootCmd.AddCommand(initCmd)
}
