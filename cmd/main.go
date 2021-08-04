package main

import (
	"fmt"
	"os"

	"github.com/kumarabd/service-template/internal/config"
	"github.com/kumarabd/service-template/pkg/service"
	"github.com/realnighthawk/bucky/logger"
)

func handleError(log logger.Handler, err error) {
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func main() {
	// Initialize Logger instance
	log, err := logger.New("service", logger.Options{
		Format: logger.SyslogLogFormat,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Config init and seed
	cfg, err := config.New()
	handleError(log, err)

	// Service Initialization
	svc, err := service.New(log, cfg)
	handleError(log, err)

	log.Info("service started")
	ch := make(chan error)
	go svc.Server.Run(ch)
	select {
	case err := <-ch:
		handleError(log, err)
	}
}
