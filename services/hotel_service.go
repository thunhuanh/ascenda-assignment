package services

import (
	"ascenda-assignment/models"
	"ascenda-assignment/repos"
	"context"
)

type HotelService interface {
	FindHotel(destination int, hotels []string) ([]*models.Hotel, error)
}

type hotelService struct {
	repo repos.HotelRepo
}

func (h *hotelService) FindHotel(destination int, hotels []string) ([]*models.Hotel, error) {
	return h.repo.GetHotel(destination, hotels)
}

func NewHotelService(ctx context.Context) *hotelService {
	return &hotelService{
		repo: repos.NewHotelRepo(ctx),
	}
}
