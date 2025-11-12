package helpers

import (
	"encoding/json"
	"os"

	"github.com/isaacwassouf/nvim-config-switcher/configs"
)

type HistoryItem struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func NewHistoryItem(from, to string) *HistoryItem {
	return &HistoryItem{
		From: from,
		To:   to,
	}
}

func AddHistoryItem(item HistoryItem) error {

	historyFilePath, err := PathFromUserCfg(configs.ToolCfgDir, configs.HistoryFile)
	if err != nil {
		return err
	}
	historyData, err := os.ReadFile(historyFilePath)

	if err != nil {
		return err
	}

	var history []HistoryItem
	if err = json.Unmarshal(historyData, &history); err != nil {
		return err
	}

	history = append(history, item)

	updatedHistoryData, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}

	if err = os.WriteFile(historyFilePath, updatedHistoryData, 0744); err != nil {
		return err
	}

	return nil
}
