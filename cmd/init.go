package cmd

import (
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

		// 1. Check if already initialized
		if hasBeenInitialized() {
			log.Fatal("The tool has already been initialized.")
		}

		toolCfgPath, err := helpers.PathFromUserCfg(configs.ToolCfgDir)
		if err != nil {
			log.Fatalf("Could not get the tool config directory: %v", err)
		}
		// 2. Create tool config directory
		err = os.MkdirAll(toolCfgPath, 0755)
		if err != nil {
			log.Fatalf("Could not create the directory: %s", toolCfgPath)
		}

		addCfgsPath, err := helpers.PathFromUserCfg(configs.ToolCfgDir, configs.AddCfgsDir)
		if err != nil {
			log.Fatalf("Could not get the tool repos directory: %v", err)
		}
		// 3. Create additional configs directory, i,e, where all the configs will be stored
		if err = os.Mkdir(addCfgsPath, 0755); err != nil {
			log.Fatalf("Could not create repos directory: %s", addCfgsPath)
		}

		// 4. Create history file
		if err = writeHistoryFile(); err != nil {
			log.Fatalf("Could not create history file: %v", err)
		}

		// 5. Move initial nvim config to tool configs directory
		if err = moveInitalNvimConfig(addCfgsPath); err != nil {
			log.Fatalf("Could not move initial nvim config: %v", err)
		}

		log.Println("Initialization completed successfully.")

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

func writeHistoryFile() error {
	// create history.json
	historyFilePath, err := helpers.PathFromUserCfg(configs.ToolCfgDir, configs.HistoryFile)
	if err != nil {
		return err
	}

	err = os.WriteFile(historyFilePath, []byte("[]"), 0744)
	if err != nil {
		return err
	}

	return nil
}

func moveInitalNvimConfig(addCfgsPath string) error {
	nvimCfgPath, err := helpers.PathFromUserCfg("nvim")
	if err != nil {
		return err
	}

	if _, err = os.Stat(nvimCfgPath); err == nil {
		destPath := path.Join(addCfgsPath, "default-nvim-config")

		err = os.Rename(nvimCfgPath, destPath)
		if err != nil {
			return err
		}

		err = os.Symlink(destPath, nvimCfgPath)
		if err != nil {
			return err
		}

		historyItem := helpers.NewHistoryItem("", "default-nvim-config")

		if err = helpers.AddHistoryItem(historyItem); err != nil {
			log.Fatalf("Could not add history item: %v", err)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}
