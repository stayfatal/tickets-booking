package models

import "ticketsbooking/libs/entities"

type RegistrationRequest struct {
	User entities.User
}

type RegistrationResponse struct {
	Error string
}

type LoginRequest struct {
	User entities.User
}

type LoginResponse struct {
	Token string
	Error string
}
