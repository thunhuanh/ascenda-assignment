package router

import (
	"ascenda-assignment/handlers"
	"ascenda-assignment/services"
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	// define services
	hotelService := services.NewHotelService(context.Background())

	// define handlers
	hotelHandler := handlers.NewHotelHandler(hotelService)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/hotels", hotelHandler.GetHotels)
	})
	return r
}
