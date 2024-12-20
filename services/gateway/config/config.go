package config

import (
	"booking/libs/utils"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func LoadConfig() (*Config, error) {
	path, err := utils.GetPath("services/gateway/config/config.yaml")
	if err != nil {
		return nil, errors.Wrap(err, "cant get path")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "cant open config file")
	}

	cfg := &Config{}

	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "cant decode config file")
	}

	return cfg, nil
}
