package main

import (
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
)

type Config struct {
	Logger *zap.Config `json:"logger"`
	//Server *routers.ServerConfig `json:"server"`
	//Db     *repository.ConfigDb  `json:"db"`
}

func NewConfig(file string) (*Config, error) {
	var cfg *Config

	rawJson, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rawJson, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
