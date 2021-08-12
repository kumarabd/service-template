package config

import (
	"os"
	"path/filepath"

	"github.com/realnighthawk/bucky/config"
	"github.com/realnighthawk/bucky/config/viper"
)

const (
	ServerConfig     = "server"
	MonitoringConfig = "monitoring"
	TracingConfig    = "tracing"
)

var (
	ApplicationName    = "default"
	ApplicationVersion = "dev"
)

// New creates a new config instance
func New() (config.Handler, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	handler, err := viper.New(viper.Options{
		FilePath: filepath.Join(wd, "internal", "config"),
		FileType: "yaml",
		FileName: "config",
	})
	if err != nil {
		return nil, err
	}

	// Seed config
	// TODO: Seed useful config

	return handler, nil
}
