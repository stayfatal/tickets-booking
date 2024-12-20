package transport

import (
	"booking/gen/authpb"
	"booking/services/auth/internal/models"
	"context"
)

func encodeRegisterResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(models.RegisterResponse)
	return &authpb.RegisterResponse{
		Token: resp.Token,
		Error: resp.Error,
	}, nil
}

func encodeLoginResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(models.LoginResponse)
	return &authpb.LoginResponse{
		Token: resp.Token,
		Error: resp.Error,
	}, nil
}
