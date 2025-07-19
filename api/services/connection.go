package services

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Connect(configPath string) error {
	cfg, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	db, err = sql.Open(
		"postgres",
		cfg.ToConnString(),
	)

	if err != nil {
		return err
	}

	db.SetMaxIdleConns(8)
	db.SetConnMaxIdleTime(3 * time.Minute)

	return db.Ping()
}
