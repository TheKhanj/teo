package main

import (
	"encoding/json"
	"os"
)

type ConfigRecord struct {
	Dir string `json:"dir"`
}

type ConfigCamera struct {
	Primary   string  `json:"primary"`
	Secondary *string `json:"secondary"`
}

type ConfigCameras map[string]ConfigCamera

type ConfigHttp struct {
	Port    *int    `json:"port"`
	Address *string `json:"address"`
}

type Config struct {
	User    *string       `json:"user"`
	Group   *string       `json:"group"`
	Http    *ConfigHttp   `json:"http"`
	Record  ConfigRecord  `json:"record"`
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
