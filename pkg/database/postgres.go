package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type ConfigDatabase struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	Username string `env:"DB_USERNAME"`
	Name     string `env:"DB_NAME"`
	Password string `env:"DB_PASSWORD"`
	SSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
}

func SetNewConnection(config ConfigDatabase) (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", config.Host, config.Port, config.Username, config.Name, config.Password, config.SSLMode),
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
