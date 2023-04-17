package handlers

import (
	"ascenda-assignment/constants"
	"ascenda-assignment/services"
	"net/http"
	"strconv"
	"strings"
)

type hotelHandler struct {
	hotelService services.HotelService
}

func (h *hotelHandler) GetHotels(w http.ResponseWriter, r *http.Request) {
	// get query params
	query := r.URL.Query()
	destinationString := query.Get("destination")
	hotelsString := query.Get("hotels")

	// validate query params
	if destinationString == "" && hotelsString == "" {
		response := Response{
			ErrorCode: 1,
			Message:   "destination and hotels are required",
			Data:      nil,
		}

		ResponseWithJson(w, http.StatusBadRequest, response)
		return
	}

	// parse the destination to int
	destination, err := strconv.Atoi(destinationString)
	if err != nil {
		response := Response{
			ErrorCode: 1,
			Message:   "Invalid query params destination",
			Data:      nil,
		}
		ResponseWithJson(w, http.StatusBadRequest, response)
		return
	}

	hotels := strings.Split(hotelsString, ",")
	data, err := h.hotelService.FindHotel(destination, hotels)
	if err == nil {
		// return in json format
		response := Response{
			ErrorCode: 0,
			Message:   "Get hotels successfully",
			Data:      data,
		}

		ResponseWithJson(w, http.StatusOK, response)
		return
	}

	if !strings.Contains(err.Error(), constants.ErrNotFound) {
		response := Response{
			ErrorCode: 1,
			Message:   "Error when get hotels " + err.Error(),
			Data:      nil,
		}
		ResponseWithJson(w, http.StatusNotFound, response)
		return
	}

	response := Response{
		ErrorCode: 1,
		Message:   "Error when get hotels " + err.Error(),
		Data:      nil,
	}
	ResponseWithJson(w, http.StatusInternalServerError, response)
}

func NewHotelHandler(hotelService services.HotelService) *hotelHandler {
	return &hotelHandler{hotelService: hotelService}
}
