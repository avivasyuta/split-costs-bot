package app

import (
	"database/sql"
	"log"
	"split-costs-bot/internal/repository/postgres"
	"split-costs-bot/internal/service"
	"split-costs-bot/pkg/database"
)

func Run() {
	db, err := database.SetNewConnection(database.ConfigDatabase{
		Host:     "localhost",
		Port:     5432,
		SSLMode:  "disable",
		Username: "split-costs-bot",
		Name:     "split-costs-bot",
		Password: "123987564",
	})
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	tripsRepo := postgres.NewTripRepository(db)
	tripsService := service.NewBookService(tripsRepo)

}
