package config

import (
	"os"
	"path/filepath"

	"github.com/realnighthawk/bucky/config"
	configprovider "github.com/realnighthawk/bucky/config/provider"
)

const (
	ServerConfig     = "server"
	MonitoringConfig = "monitoring"
)

// New creates a new config instance
func New() (config.Handler, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	handler, err := configprovider.NewViper(configprovider.Options{
		FilePath: filepath.Join(wd, "internal", "config"),
		FileType: "yaml",
		FileName: "config",
	})
	if err != nil {
		return nil, err
	}

	return handler, nil
}
