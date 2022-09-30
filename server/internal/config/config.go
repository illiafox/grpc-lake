package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

func NewConfig() (Config, error) {
	cfg := Config{
		Flags: ParseFlags(),
	}

	//err := cleanenv.ReadConfig(cfg.Flags.ConfigPath, &cfg)
	//if err != nil {
	//	return Config{}, fmt.Errorf("yaml: %w", err)
	//}

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("env: %w", err)
	}

	return cfg, nil
}
