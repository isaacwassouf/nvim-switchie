package cmd

import (
	"log"
	"os"
	"path"

	"github.com/isaacwassouf/nvim-config-switcher/helpers"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:  "use",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		installedCfgsPath, err := helpers.GetInstalledCfgsPath()
		if err != nil {
			log.Fatal(err)
		}

		cfgPath := path.Join(installedCfgsPath, name)
		if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
			log.Fatalf("Repo with name %s does not exist", name)
		}

		// check if a config currently exists
		currentNvimCfg, err := helpers.GetNvimCfgPath()
		if err != nil {
			log.Fatal(err)
		}
		if _, err := os.Stat(currentNvimCfg); err == nil {
			os.Rename(currentNvimCfg, currentNvimCfg+".bak")
		}

		err = os.Symlink(cfgPath, currentNvimCfg)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Switched to repo %s", name)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
