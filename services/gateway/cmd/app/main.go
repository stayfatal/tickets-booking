package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"ticketsbooking/libs/config"
	"ticketsbooking/libs/log"
	"ticketsbooking/services/gateway/internal/service"
	transport "ticketsbooking/services/gateway/internal/transport/http"
)

func main() {
	logger := log.New()

	cfg, err := config.LoadConfigs()
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}

	svc, err := service.New(cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	defer svc.GratefulStop()

	srv := transport.NewGatewayServer(svc, logger)

	quit := make(chan struct{})
	go func() {
		logger.Info().Msgf("Server is now listening on port: %s", cfg.Gateway.Port)
		if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Gateway.Port), srv); err != nil {
			logger.Error().Err(err).Msg("")
			quit <- struct{}{}
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		sig := <-c
		logger.Info().Msg(sig.String())
		quit <- struct{}{}
	}()

	<-quit
}
