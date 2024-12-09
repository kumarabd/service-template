package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kumarabd/service-template/pkg/service"
	"github.com/kumarabd/service-template/internal/logger"
	"github.com/kumarabd/service-template/internal/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HTTPServerConfig struct {
	Port string `json:"port" yaml:"port"`
}

type HTTPServer struct {
	handler *gin.Engine
	service *service.Handler
	log     *logger.Handler
	metric  *metrics.Handler
}

func (h *HTTPServer) MetricsHandler(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

func (h *HTTPServer) HealthHandler(c *gin.Context) {
	c.JSON(200, http.StatusText(http.StatusOK))
}
