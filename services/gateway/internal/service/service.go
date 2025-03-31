package service

import (
	"context"
	"fmt"
	"ticketsbooking/gen/authpb"
	"ticketsbooking/libs/config"
	"ticketsbooking/libs/entities"
	"ticketsbooking/services/gateway/internal/interfaces"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type service struct {
	conns   []*grpc.ClientConn
	authSvc authpb.AuthenticationClient
}

func New(cfg *config.ServicesConfig) (interfaces.Service, error) {
	conns := []*grpc.ClientConn{}

	clientConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Auth.Host, cfg.Auth.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	authSvc := authpb.NewAuthenticationClient(clientConn)
	conns = append(conns, clientConn)

	return &service{conns: conns, authSvc: authSvc}, nil
}

func (s *service) Register(user entities.User) (*authpb.RegisterResponse, error) {
	return s.authSvc.Register(context.Background(), &authpb.RegisterRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
}

func (s *service) Login(user entities.User) (*authpb.LoginResponse, error) {
	return s.authSvc.Login(context.Background(), &authpb.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	})
}

func (s *service) GratefulStop() {
	for _, conn := range s.conns {
		conn.Close()
	}
}
