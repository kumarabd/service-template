package service

import (
	internal_config "github.com/kumarabd/service-template/internal/config"
	"github.com/realnighthawk/bucky/apm"
	cache_package "github.com/realnighthawk/bucky/cache"
	config_package "github.com/realnighthawk/bucky/config"
	"github.com/realnighthawk/bucky/logger"
	"github.com/realnighthawk/bucky/server"
	"github.com/realnighthawk/bucky/server/http"
	"github.com/realnighthawk/bucky/tracing"
)

type Handler struct {
	Server server.Server
	log    logger.Handler
	config config_package.Handler
	cache  cache_package.Handler
	//database
	//broker
}

func New(l logger.Handler, config config_package.Handler, cache cache_package.Handler) (*Handler, error) {
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

	// Enable Metrics for this application
	if config.Is(internal_config.MonitoringConfig) {
		mOpts := apm.Options{}
		err := config.GetObject(internal_config.MonitoringConfig, &mOpts)
		if err != nil {
			return nil, err
		}
		s.EnableMetrics(mOpts)
	}

	// Enable Tracing for this application
	if config.Is(internal_config.TracingConfig) {
		tOpts := tracing.Options{}
		err := config.GetObject(internal_config.MonitoringConfig, &tOpts)
		if err != nil {
			return nil, err
		}
		// TODO: Add tracing init
	}

	return &Handler{
		log:    l,
		Server: s,
		config: config,
		cache:  cache,
	}, nil
}
