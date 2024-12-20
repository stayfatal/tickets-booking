package main

import (
	"booking/libs/log"
	"booking/services/gateway/config"
	"booking/services/gateway/internal/service"
	transport "booking/services/gateway/internal/transport/http"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := log.New()

	cfg, err := config.LoadConfig()
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
		logger.Info().Msgf("Server is now listening on port: %s", cfg.Server.Port)
		if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.Port), srv); err != nil {
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
