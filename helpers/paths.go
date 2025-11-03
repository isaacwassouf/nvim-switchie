package helpers

import (
	"os"
	"path"
)

func GetBaseCfgPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, ".config"), nil
}

func GetInstalledCfgsPath() (string, error) {
	baseCfgPath, err := GetBaseCfgPath()
	if err != nil {
		return "", err
	}
	return path.Join(baseCfgPath, "switchie", "repos"), nil
}

func GetNvimCfgPath() (string, error) {
	baseCfgPath, err := GetBaseCfgPath()
	if err != nil {
		return "", err
	}
	return path.Join(baseCfgPath, "nvim"), nil
}
