package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Environment string `json:"environment"`
	TestCommand string `json:"testCommand"`
	Auth        struct {
		Provider    string `json:"provider"`
		SessionRole string `json:"sessionRole"`
		SessionDir  string `json:"sessionDir"`
	} `json:"auth"`
}

func Load(path string) (Config, error) {
	var cfg Config
	b, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("read config: %w", err)
	}
	if err := json.Unmarshal(b, &cfg); err != nil {
		return cfg, fmt.Errorf("parse config: %w", err)
	}
	if cfg.Auth.SessionDir == "" {
		cfg.Auth.SessionDir = ".auth"
	}
	return cfg, nil
}
