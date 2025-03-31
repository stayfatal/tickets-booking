package models

import "ticketsbooking/libs/entities"

type RegisterRequest struct {
	User entities.User
}

type RegisterResponse struct {
	Error string
}

type LoginRequest struct {
	User entities.User
}

type LoginResponse struct {
	Token string
	Error string
}
