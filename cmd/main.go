package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kumarabd/service-template/internal/config"
	"github.com/kumarabd/service-template/internal/logger"
	"github.com/kumarabd/service-template/internal/metrics"
	"github.com/kumarabd/service-template/pkg/server"
	"github.com/kumarabd/service-template/pkg/service"
)

// main is the entry point of the application
func main() {
	// Initialize a new logger with the application name and syslog format
	log, err := logger.New(config.ApplicationName, logger.Options{
		Format: logger.SyslogLogFormat,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Initialize a new configuration handler
	configHandler, err := config.New()
	if err != nil {
		log.Error().Err(err).Msg("")
		os.Exit(1)
	}

	// Initialize a new metrics handler with the application name and server namespace
	metricsHandler, err := metrics.New(config.ApplicationName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Initialize a new service with the logger, metrics handler, data layer, and service configuration
	service, err := service.New(log, metricsHandler, nil, configHandler.Service)
	if err != nil {
		log.Error().Err(err).Msg("service initialization failed")
		os.Exit(1)
	}
	log.Info().Msg("service initialized")

	// Initialize a new server with the logger, metrics handler, server configuration, and service
	srv, err := server.New(config.ApplicationName, log, metricsHandler, configHandler.Server, service)
	if err != nil {
		log.Error().Err(err).Msg("")
		os.Exit(1)
	}
	log.Info().Msg("server initialized")

	// Create a channel to control the server
	ch := make(chan struct{})

	// Run the server
	log.Info().Msg("server starting")
	srv.Start(ch)
	log.Info().Msg("server running")

	// Create a signal channel to handle graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a stop signal or service/server completion
	exit := false
	for !exit {
		select {
		case <-ch:
			exit = true
		case <-signalChan:
			exit = true
		}
	}

	log.Info().Msg("received stop. gracefully shutting down...")
	close(ch)
}
