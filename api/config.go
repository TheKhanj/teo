package main

import (
	"encoding/json"
	"os"
)

type ConfigCamera struct {
	Primary   string  `json:"primary"`
	Secondary *string `json:"secondary"`
}

type ConfigCameras map[string]ConfigCamera

type ConfigApi struct {
	Port    *int    `json:"port"`
	Address *string `json:"address"`
}

type Config struct {
	Api     *ConfigApi    `json:"api"`
	Cameras ConfigCameras `json:"cameras"`
}

func parseConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
