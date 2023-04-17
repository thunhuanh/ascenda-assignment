package repos

import (
	"ascenda-assignment/constants"
	"ascenda-assignment/models"

	"context"
	"strings"
	"testing"
)

var testHotelRepo HotelRepo

func TestMain(m *testing.M) {
	// setup repo
	testHotelRepo = NewHotelRepo(context.Background())
	m.Run()
}

func insertSampleHotel(isMultiple bool) ([]*models.Hotel, error) {
	if !isMultiple {
		sample := &models.Hotel{
			HotelID:       "test",
			DestinationID: 1234,
			Name:          "test",
		}
		return []*models.Hotel{sample}, testHotelRepo.InsertHotel(*sample)
	}
	samples := []*models.Hotel{
		{
			HotelID:       "test1",
			DestinationID: 1234,
			Name:          "test1",
		},
		{
			HotelID:       "test2",
			DestinationID: 1234,
			Name:          "test2",
		},
	}
	err := testHotelRepo.InsertManyHotel(samples)
	return samples, err
}

func Test_hotelRepo_InsertHotel(t *testing.T) {
	// first, we need to create a sample hotel and insert it into the database
	// to test the InsertHotel function we just need to query the database for the sample hotel
	// and check if the result is the same as the sample hotel

	sample, err := insertSampleHotel(false)
	// clean up the database
	defer func() {
		for _, sample := range sample {
			_ = testHotelRepo.DeleteHotel(sample.HotelID)
		}
	}()
	if err != nil && !strings.Contains(err.Error(), constants.ErrDuplicateKey) {
		t.Errorf("Failed insert test: %v", err)
		t.FailNow()
	}

	// query the database for the sample hotel
	hotel, err := testHotelRepo.GetHotel(sample[0].DestinationID, []string{sample[0].HotelID})
	if err != nil {
		// t.Errorf("Failed insert test: %v", err)
		t.Log(hotel)
		t.FailNow()
	}

	// check if the result is the same as the sample hotel
	if !isHotelEqual(sample[0], hotel[0]) {
		t.Errorf("Failed insert test: %v", err)
		t.FailNow()
	}
}

func Test_hotelRepo_InsertManyHotel(t *testing.T) {
	// first, we need to create a sample hotel and insert it into the database
	// to test the InsertHotel function we just need to query the database for the sample hotel
	// and check if the result is the same as the sample hotel

	samples, err := insertSampleHotel(true)
	// clean up the database
	defer func() {
		for _, sample := range samples {
			_ = testHotelRepo.DeleteHotel(sample.HotelID)
		}
	}()
	if err != nil && !strings.Contains(err.Error(), constants.ErrDuplicateKey) {
		t.Errorf("Failed insert test: %v", err)
		t.FailNow()
	}

	for _, sample := range samples {
		// query the database for the sample hotel
		hotel, err := testHotelRepo.GetHotel(sample.DestinationID, []string{sample.HotelID})
		if err != nil {
			t.Errorf("Failed insert test: %v", err)
			t.FailNow()
		}

		// check if the result is the same as the sample hotel
		if !isHotelEqual(sample, hotel[0]) {
			t.Errorf("Failed insert test: %v", err)
			t.FailNow()
		}
	}

}

func isHotelEqual(h1, h2 *models.Hotel) bool {
	if h1.HotelID != h2.HotelID {
		return false
	}
	if h1.DestinationID != h2.DestinationID {
		return false
	}
	if h1.Name != h2.Name {
		return false
	}
	return true
}
