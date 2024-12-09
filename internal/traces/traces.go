package traces

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Handler struct {
}

type Options struct {
}

func New(name string, namespace string) (*Handler, error) {
	return &Handler{}, nil
}

func NewGinMiddleware(name string) gin.HandlerFunc {
	return otelgin.Middleware(name)
}
