package service

import (
	internal_config "github.com/kumarabd/service-template/internal/config"
	"github.com/realnighthawk/bucky/apm/prometheus"
	cache_package "github.com/realnighthawk/bucky/cache"
	"github.com/realnighthawk/bucky/logger"
	"github.com/realnighthawk/bucky/server"
	"github.com/realnighthawk/bucky/server/http/gin"
)

type Handler struct {
	Server server.Server
	log    logger.Handler
	cache  cache_package.Handler
	//database
	//broker
}

func New(l logger.Handler, config *internal_config.Static, cache cache_package.Handler) (*Handler, error) {

	// Initiate Server object
	s, err := gin.New(config.Server)
	if err != nil {
		return nil, err
	}

	h := &Handler{
		log:    l,
		Server: s,
		cache:  cache,
	}

	// Enable Metrics for this application
	if config.Monitoring.Prometheus.Enabled {
		h.registerMetrics()
	}

	// Enable Tracing for this application
	if config.Tracing.OTel.Enabled {
		// TODO: Add tracing init
	}

	// Seed HTTP routes
	h.registerRoutes()

	return h, nil
}

func (h *Handler) registerMetrics() {
	h.Server.(*gin.Server).RegisterGenericWithGroup("monitoring", "GET", "/prometheus", prometheus.GetHTTPHandler())
}

func (h *Handler) registerRoutes() {
	h.Server.(*gin.Server).RegisterWithGroup("application", "GET", "/status", h.statusHandler)
}
