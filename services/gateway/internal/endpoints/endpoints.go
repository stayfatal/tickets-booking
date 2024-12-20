package endpoints

import (
	"booking/services/gateway/internal/interfaces"
	"booking/services/gateway/internal/models"
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Register endpoint.Endpoint
	Login    endpoint.Endpoint
}

func MakeEndpoints(svc interfaces.Service) *Endpoints {
	return &Endpoints{
		Register: makeRegisterEndpoint(svc),
		Login:    makeLoginEndpoint(svc),
	}
}

func makeRegisterEndpoint(svc interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(models.RegistrationRequest)
		if !ok {
			return nil, errors.New("type assertion error")
		}
		return svc.Register(req.User)
	}
}

func makeLoginEndpoint(svc interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(models.LoginRequest)
		if !ok {
			return nil, errors.New("type assertion error")
		}
		return svc.Login(req.User)
	}
}
