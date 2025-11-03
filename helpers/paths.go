package helpers

import (
	"os"
	"path"
)

func GetBaseConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, ".config"), nil
}

func GetReposPath() (string, error) {
	baseCfgPath, err := GetBaseConfigPath()
	if err != nil {
		return "", err
	}
	return path.Join(baseCfgPath, "switchie", "repos"), nil
}

func GetNvimConfigPath() (string, error) {
	baseCfgPath, err := GetBaseConfigPath()
	if err != nil {
		return "", err
	}
	return path.Join(baseCfgPath, "nvim"), nil
}
