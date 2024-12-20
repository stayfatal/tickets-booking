package transport

import (
	"booking/gen/authpb"
	"booking/services/gateway/internal/models"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func encodeRegisterResponse(_ context.Context, w http.ResponseWriter, i interface{}) error {
	resp, ok := i.(*authpb.RegisterResponse)
	if !ok {
		return errors.New("type assertion error")
	}
	response := models.RegistrationResponse{
		Token: resp.Token,
	}
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func encodeLoginResponse(_ context.Context, w http.ResponseWriter, i interface{}) error {
	resp, ok := i.(*authpb.LoginResponse)
	if !ok {
		return errors.New("type assertion error")
	}
	response := models.LoginResponse{
		Token: resp.Token,
	}
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
