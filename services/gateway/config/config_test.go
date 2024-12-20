package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, cfg)
	assert.NotNil(t, cfg.Server)
	assert.NotNil(t, cfg.Services)
}
