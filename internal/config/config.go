package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/realnighthawk/bucky/config"
	configprovider "github.com/realnighthawk/bucky/config/provider"
)

const (
	ServerConfig     = "server"
	MonitoringConfig = "monitoring"
	TracingConfig    = "tracing"
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

	if os.Getenv("SERVER_VERSION") != "" {
		handler.SetKey(fmt.Sprintf("%s.version", ServerConfig), os.Getenv("SERVER_VERSION"))
	}

	return handler, nil
}
