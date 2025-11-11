package cmd

import (
	"log"
	"os"
	"path"

	"github.com/isaacwassouf/nvim-config-switcher/configs"
	"github.com/isaacwassouf/nvim-config-switcher/helpers"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:  "use",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		addCfgPath, err := helpers.PathFromUserCfg(configs.ToolCfgDir, configs.AddCfgsDir)
		if err != nil {
			log.Fatal(err)
		}

		// check if the config with the given name exists
		cfgPath := path.Join(addCfgPath, name)
		if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
			log.Fatalf("Repo with name %s does not exist", name)
		}

		currentNvimCfg, err := helpers.PathFromUserCfg("nvim")
		if err != nil {
			log.Fatal(err)
		}

		// if there's a backup already, remove it
		if _, err = os.Stat(currentNvimCfg + ".bak"); err == nil {
			if err = os.RemoveAll(currentNvimCfg + ".bak"); err != nil {
				log.Fatal(err)
			}
		}

		// rename current config to backup
		if _, err = os.Stat(currentNvimCfg); err == nil {
			if err = os.Rename(currentNvimCfg, currentNvimCfg+".bak"); err != nil {
				log.Fatal(err)
			}
		}

		// create a symlink to the new config
		if err = os.Symlink(cfgPath, currentNvimCfg); err != nil {
			log.Fatal(err)
		}

		log.Printf("Switched to repo %s", name)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
