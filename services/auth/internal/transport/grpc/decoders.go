package transport

import (
	"booking/gen/authpb"
	"booking/libs/entities"
	"booking/services/auth/internal/models"
	"context"
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
