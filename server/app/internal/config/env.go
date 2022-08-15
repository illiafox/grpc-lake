package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

func New() (Config, error) {
	var conf Config

	err := cleanenv.ReadEnv(&conf)
	if err != nil {
		return Config{}, fmt.Errorf("read env: %w", err)
	}

	return conf, nil
}
