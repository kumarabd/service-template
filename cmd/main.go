package main

import (
	"fmt"
	"os"

	"github.com/kumarabd/service-template/internal/cache"
	"github.com/kumarabd/service-template/internal/channels"
	"github.com/kumarabd/service-template/internal/config"
	"github.com/kumarabd/service-template/pkg/service"
	"github.com/realnighthawk/bucky/logger"
)

func main() {
	// Initialize Logger instance
	log, err := logger.New(config.ApplicationName, logger.Options{
		Format: logger.SyslogLogFormat,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Config init and seed
	configHandler, err := config.New()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Cache init and seed
	cacheHandler, err := cache.New()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Service Initialization
	svc, err := service.New(log, configHandler, cacheHandler)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Info("service started")
	ch := channels.NewServerChannel()
	go svc.Server.Run(ch)
	select {
	case err := <-ch:
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
	}
}
