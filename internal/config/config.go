package config

import (
	"os"
	"path/filepath"

	"github.com/realnighthawk/bucky/apm"
	"github.com/realnighthawk/bucky/config/viper"
	"github.com/realnighthawk/bucky/server/http/gin"
	"github.com/realnighthawk/bucky/tracing"
)

var (
	ApplicationName    = "default"
	ApplicationVersion = "dev"
)

type Static struct {
	Server     gin.Options     `json:"server,omitempty" yaml:"server,omitempty"`
	Monitoring apm.Options     `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`
	Tracing    tracing.Options `json:"tracing,omitempty" yaml:"tracing,omitempty"`
}

// New creates a new config instance
func New() (*Static, error) {
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
	static := &Static{}
	handler.GetAll(&static)

	return static, nil
}
