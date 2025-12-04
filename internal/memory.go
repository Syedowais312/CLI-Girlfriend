package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"github.com/Syedowais312/CLI-Girlfriend/model"
)

func getHistoryPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".my-girlfriend")

	// create folder if missing
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	return filepath.Join(dir, "history.json"), nil
}

func LoadHistory() (*model.ChatHistory, error) {
	path, err := getHistoryPath()
	if err != nil {
		return &model.ChatHistory{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return &model.ChatHistory{}, nil
	}

	var history model.ChatHistory
	if err := json.Unmarshal(data, &history); err != nil {
		return &model.ChatHistory{}, nil
	}

	return &history, nil
}


// save the memory to disk
func SaveHistory(history *model.ChatHistory) error {
	path, err := getHistoryPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}


func ClearHistory() error {
	path, err := getHistoryPath()
	if err != nil {
		return err
	}


	
	// overwrite with empty chat history
	emptyHistory := model.ChatHistory{Messages: []model.ChatMessage{}}

	data, err := json.MarshalIndent(emptyHistory, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

