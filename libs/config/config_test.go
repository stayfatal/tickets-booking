package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	test := map[string]string{
		"AUTH_HOST":    "auth",
		"AUTH_PORT":    "8020",
		"GATEWAY_HOST": "gateway",
		"GATEWAY_PORT": "8030",
		"BOOKING_HOST": "booking",
		"BOOKING_PORT": "8040",
	}

	for key, val := range test {
		os.Setenv(key, val)
	}

	cfg, err := LoadConfigs()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "auth", cfg.Auth.Host)
	assert.Equal(t, "8020", cfg.Auth.Port)

	assert.Equal(t, "gateway", cfg.Gateway.Host)
	assert.Equal(t, "8030", cfg.Gateway.Port)

	assert.Equal(t, "booking", cfg.Booking.Host)
	assert.Equal(t, "8040", cfg.Booking.Port)
}
