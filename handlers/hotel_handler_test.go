package handlers

import (
	"ascenda-assignment/models"
	"ascenda-assignment/services"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_hotelHandler_GetHotels(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/hotels?destination=5432", nil)
	w := httptest.NewRecorder()

	// define hotel servies
	hotelService := services.NewHotelService(context.Background())
	handler := NewHotelHandler(hotelService)
	handler.GetHotels(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	result := Response{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	// for simplicity, we only check the length of the data
	expectedData := make([]models.Hotel, 2)
	if result.Data == nil {
		t.Errorf("expected data to be %v got %v", expectedData, result.Data)
		t.FailNow()
	} else if len(result.Data.([]interface{})) != len(expectedData) {
		t.Errorf("expected data to be %v got %v", expectedData, result.Data)
	}

}
