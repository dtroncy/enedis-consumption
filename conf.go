package enedisconsumption

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token string `json:"token"`
	PRM   string `json:"prm"`
}

func loadConfig(filePath string) (Config, error) {
	var config Config

	file, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
