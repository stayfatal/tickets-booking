package service

import (
	"fmt"
	"testing"
	"ticketsbooking/libs/config"
	"ticketsbooking/libs/entities"
	"ticketsbooking/libs/publicauth"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAndLogin(t *testing.T) {
	cfg, err := config.LoadConfigs()
	if err != nil {
		t.Fatal(err)
	}

	svc, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	expected := entities.User{
		Name:     "test",
		Email:    fmt.Sprintf("test%s@gmail.com", uuid.New().String()),
		Password: "123",
	}

	regResp, err := svc.Register(expected)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, regResp)

	loginResp, err := svc.Login(expected)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, loginResp)

	loginClaims, err := publicauth.ValidateToken(loginResp.Token)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, loginClaims)
	assert.Equal(t, expected.Email, loginClaims.Email)
}
