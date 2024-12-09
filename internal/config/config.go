package config

import (
	"os"

	"github.com/kumarabd/service-template/internal/metrics"
	"github.com/kumarabd/service-template/pkg/server"
	"github.com/kumarabd/service-template/pkg/service"
	"gopkg.in/yaml.v2"
)

var (
	ApplicationName    = "default"
	ApplicationVersion = "dev"
)

type Config struct {
	Server  *server.Config   `json:"server,omitempty" yaml:"server,omitempty"`
	Service *service.Config  `json:"service" yaml:"service"`
	Metrics *metrics.Options `json:"metrics,omitempty" yaml:"metrics,omitempty"`
	//Traces  *traces.Options  `json:"traces,omitempty" yaml:"traces,omitempty"`
}

// New creates a new config instance
func New() (*Config, error) {
	configFilePath := os.Getenv("CONFIG_PATH")
	if len(configFilePath) == 0 {
		configFilePath = "/app/config.yaml"
	}

	configContent, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	configObject := Config{}
	err = yaml.Unmarshal(configContent, &configObject)
	if err != nil {
		return nil, err
	}
	return &configObject, nil
}
