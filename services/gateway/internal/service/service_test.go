package service

import (
	"booking/libs/entities"
	"booking/libs/publicauth"
	"booking/services/gateway/config"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAndLogin(t *testing.T) {
	cfg, err := config.LoadConfig()
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

	regClaims, err := publicauth.ValidateToken(regResp.Token)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, regClaims)
	assert.Equal(t, expected.Email, regClaims.Email)

	loginResp, err := svc.Login(expected)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, loginResp)

	loginClaims, err := publicauth.ValidateToken(loginResp.Token)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, regClaims)
	assert.Equal(t, expected.Email, loginClaims.Email)
}
