package services

import (
	"encoding/json"
	"fmt"
	"os"
)

type databaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (cfg databaseConfig) ToConnString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
}

func loadConfig(configPath string) (*databaseConfig, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	cfg := &databaseConfig{}
	err = json.NewDecoder(f).Decode(cfg)
	return cfg, err
}
