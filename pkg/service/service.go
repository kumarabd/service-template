package service

import (
	internal_config "github.com/kumarabd/service-template/internal/config"
	"github.com/realnighthawk/bucky/apm"
	config_package "github.com/realnighthawk/bucky/config"
	"github.com/realnighthawk/bucky/logger"
	"github.com/realnighthawk/bucky/server"
	"github.com/realnighthawk/bucky/server/http"
)

type Handler struct {
	Server server.Server
	log    logger.Handler
	config config_package.Handler
	cache  config_package.Handler
	//database
	//broker
}

func New(l logger.Handler, config config_package.Handler) (*Handler, error) {
	// Initiate Server object
	sOpts := server.Options{}
	err := config.GetObject(internal_config.ServerConfig, &sOpts)
	if err != nil {
		return nil, err
	}
	s, err := http.New(sOpts)
	if err != nil {
		return nil, err
	}

	// Enable Prometheus metrics for this application
	mOpts := apm.Options{}
	err = config.GetObject(internal_config.MonitoringConfig, &mOpts)
	if err != nil {
		return nil, err
	}
	s.EnableMetrics(mOpts)

	return &Handler{
		log:    l,
		Server: s,
		config: config,
	}, nil
}
