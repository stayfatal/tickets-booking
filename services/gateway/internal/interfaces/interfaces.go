package interfaces

import (
	"booking/gen/authpb"
	"booking/libs/entities"
)

type Service interface {
	Register(user entities.User) (*authpb.RegisterResponse, error)
	Login(user entities.User) (*authpb.LoginResponse, error)
	GratefulStop()
}
