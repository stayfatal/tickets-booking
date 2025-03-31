package transport

import (
	"context"
	"ticketsbooking/gen/authpb"
	"ticketsbooking/libs/entities"
	"ticketsbooking/services/auth/internal/models"
)

func decodeRegisterRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*authpb.RegisterRequest)
	return models.RegisterRequest{
		User: entities.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		},
	}, nil
}

func decodeLoginRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*authpb.LoginRequest)
	return models.LoginRequest{
		User: entities.User{
			Email:    req.Email,
			Password: req.Password,
		},
	}, nil
}
