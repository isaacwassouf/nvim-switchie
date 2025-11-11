package helpers

import (
	"os"
	"path"
)

func GetUserCfgPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, ".config"), nil
}

func GetToolCfgPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, ".config", "switchie"), nil
}

func GetInstalledCfgsPath() (string, error) {
	toolCfgPath, err := GetToolCfgPath()
	if err != nil {
		return "", err
	}
	return path.Join(toolCfgPath, "repos"), nil
}

func GetNvimCfgPath() (string, error) {
	userCfgPath, err := GetUserCfgPath()
	if err != nil {
		return "", err
	}
	return path.Join(userCfgPath, "nvim"), nil
}
