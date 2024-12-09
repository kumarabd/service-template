package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Handler struct {
	RequestsReceived *prometheus.CounterVec
}

type Options struct {
	// Additional labels necessary
}

func New(name string) (*Handler, error) {
	return &Handler{
		RequestsReceived: promauto.NewCounterVec(prometheus.CounterOpts{
			Name:        "http_requests_received",
			ConstLabels: map[string]string{
				// Add labels
			},
			Help: "The total number of http requests received",
		}, []string{"status"}),
	}, nil
}
