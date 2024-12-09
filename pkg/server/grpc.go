package server

import (
	"github.com/kumarabd/service-template/internal/logger"
	"github.com/kumarabd/service-template/internal/metrics"
	"github.com/kumarabd/service-template/pkg/service"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	// register methods
	handler *grpc.Server
	service *service.Handler
	log     *logger.Handler
	metric  *metrics.Handler
}

type GRPCServerConfig struct {
	Port string `json:"port" yaml:"port"`
}
