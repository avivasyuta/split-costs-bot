package postgres

import (
	"database/sql"
	"split-costs-bot/internal/domain"
)

type TripRepository struct {
	db *sql.DB
}

func NewTripRepository(db *sql.DB) *TripRepository {
	return &TripRepository{db}
}

func (repo *TripRepository) Create(trip domain.Trip) (domain.Trip, error) {
	query := "INSERT INTO trips (name) values ($1) RETURNING id"
	err := repo.db.QueryRow(query, trip.Name).Scan(&trip.ID)

	return trip, err
}
