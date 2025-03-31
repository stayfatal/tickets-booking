package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"ticketsbooking/gen/authpb"
	"ticketsbooking/services/gateway/internal/models"

	"github.com/pkg/errors"
)

func encodeRegisterResponse(_ context.Context, w http.ResponseWriter, i interface{}) error {
	resp, ok := i.(*authpb.RegisterResponse)
	if !ok {
		return errors.New("type assertion error")
	}
	response := models.RegistrationResponse{
		Error: resp.Error,
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
		Error: resp.Error,
		Token: resp.Token,
	}
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
