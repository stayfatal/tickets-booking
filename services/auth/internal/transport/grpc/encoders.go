package transport

import (
	"context"
	"ticketsbooking/gen/authpb"
	"ticketsbooking/services/auth/internal/models"
)

func encodeRegisterResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(models.RegisterResponse)
	return &authpb.RegisterResponse{
		Error: resp.Error,
	}, nil
}

func encodeLoginResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(models.LoginResponse)
	return &authpb.LoginResponse{
		Error: resp.Error,
		Token: resp.Token,
	}, nil
}
