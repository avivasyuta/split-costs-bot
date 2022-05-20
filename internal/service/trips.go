package service

import (
	"split-costs-bot/internal/domain"
	"split-costs-bot/internal/repository/postgres"
)

type TripsService struct {
	repo *postgres.TripRepository
}

func NewBookService(repo *postgres.TripRepository) *TripsService {
	return &TripsService{repo}
}

func (service *TripsService) Create(input domain.TripInput) (domain.Trip, error) {
	trip := domain.Trip{}

	if input.Name != nil {
		trip.Name = *input.Name
	}

	return service.repo.Create(trip)
}
