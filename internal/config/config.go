package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type HttpServer struct {
	Address string `yaml:"address"`
}

type Config struct {
	HttpServer HttpServer `yaml:"http_server"`
}

func NewConfig(filename string) (*Config, error) {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
