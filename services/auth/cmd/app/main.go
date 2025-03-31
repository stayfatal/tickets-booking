package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"ticketsbooking/gen/authpb"
	"ticketsbooking/libs/config"
	"ticketsbooking/libs/log"
	"ticketsbooking/services/auth/internal/cache"
	"ticketsbooking/services/auth/internal/repository"
	"ticketsbooking/services/auth/internal/service"
	transport "ticketsbooking/services/auth/internal/transport/grpc"

	"google.golang.org/grpc"
)

func main() {
	log := log.New()

	cfg, err := config.LoadConfigs()
	if err != nil {
		log.Fatal().Err(err).Msg("failed loading cfg")
	}

	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed opening db")
	}
	defer db.Close()

	cacheDB, err := config.NewRedisDB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed opening cache db")
	}

	repo := repository.New(db)

	cache := cache.New(cacheDB)

	svc := service.New(repo, cache)

	authServer := transport.NewGRPCServer(svc, log)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Auth.Port))
	if err != nil {
		log.Fatal().Err(err).Msg("failed starting listener")
	}

	srv := grpc.NewServer()
	defer srv.GracefulStop()

	authpb.RegisterAuthenticationServer(srv, authServer)

	exit := make(chan struct{})
	go func() {
		log.Info().Msgf("Server is now listening on port: %s", cfg.Auth.Port)
		if err := srv.Serve(l); err != nil {
			log.Error().Err(err).Msg("")
			exit <- struct{}{}
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		sig := <-c
		log.Info().Msg(sig.String())
		exit <- struct{}{}
	}()

	<-exit
}
