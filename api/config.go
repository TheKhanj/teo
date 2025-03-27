package main

import (
	"encoding/json"
	"os"
)

type ConfigPreset struct {
	Stream     string   `json:"stream"`
	Fps        *float32 `json:"fps"`
	Resolution *string  `json:"resolution"`
}

type ConfigCamera struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}

type ConfigCameras map[string]ConfigCamera

type ConfigApi struct {
	Port                         *int                     `json:"port"`
	Address                      *string                  `json:"address"`
	DefaultNonActiveCameraPreset *string                  `json:"defaultNonActiveCameraPreset"`
	DefaultActiveCameraPreset    *string                  `json:"defaultActiveCameraPreset"`
	Presets                      *map[string]ConfigPreset `json:"presets"`
}

type User struct {
	Password string `json:"password"`
}

type Config struct {
	Api     *ConfigApi       `json:"api"`
	Users   *map[string]User `json"users"`
	Cameras ConfigCameras    `json:"cameras"`
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

	if config.Api == nil {
		config.Api = &ConfigApi{}
	}
	if config.Api.Port == nil {
		defaultPort := 8081
		config.Api.Port = &defaultPort
	}
	if config.Api.Address == nil {
		defaultAddress := "0.0.0.0"
		config.Api.Address = &defaultAddress
	}
	if config.Api.DefaultActiveCameraPreset == nil {
		defaultActiveCameraPreset := "primary"
		config.Api.DefaultActiveCameraPreset = &defaultActiveCameraPreset
	}
	if config.Api.DefaultNonActiveCameraPreset == nil {
		defaultNonActiveCameraPreset := "primary"
		config.Api.DefaultNonActiveCameraPreset = &defaultNonActiveCameraPreset
	}
	if config.Api.Presets == nil {
		defaultPresets := make(map[string]ConfigPreset)
		config.Api.Presets = &defaultPresets
	}

	return config, nil
}
