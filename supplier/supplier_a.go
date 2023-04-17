package supplier

import (
	"ascenda-assignment/models"
	"encoding/json"
	"strconv"
)

/*
supplier A data: [
  {
    "Id": "iJhz",
    "DestinationId": 5432,
    "Name": "Beach Villas Singapore",
    "Latitude": 1.264751,
    "Longitude": 103.824006,
    "Address": " 8 Sentosa Gateway, Beach Villas ",
    "City": "Singapore",
    "Country": "SG",
    "PostalCode": "098269",
    "Description": "  This 5 star hotel is located on the coastline of Singapore.",
    "Facilities": ["Pool", "BusinessCenter", "WiFi ", "DryCleaning", " Breakfast"]
  },
  {
    "Id": "SjyX",
    "DestinationId": 5432,
    "Name": "InterContinental Singapore Robertson Quay",
    "Latitude": null,
    "Longitude": null,
    "Address": " 1 Nanson Road",
    "City": "Singapore",
    "Country": "SG",
    "PostalCode": "238909",
    "Description": "Enjoy sophisticated waterfront living at the new InterContinental® Singapore Robertson Quay, luxury's preferred address nestled in the heart of Robertson Quay along the Singapore River, with the CBD just five minutes drive away. Magnifying the comforts of home, each of our 225 studios and suites features a host of thoughtful amenities that combine modernity with elegance, whilst maintaining functional practicality. The hotel also features a chic, luxurious Club InterContinental Lounge.",
    "Facilities": ["Pool", "WiFi ", "Aircon", "BusinessCenter", "BathTub", "Breakfast", "DryCleaning", "Bar"]
  },
  {
    "Id": "f8c9",
    "DestinationId": 1122,
    "Name": "Hilton Shinjuku Tokyo",
    "Latitude": "",
    "Longitude": "",
    "Address": "160-0023, SHINJUKU-KU, 6-6-2 NISHI-SHINJUKU, JAPAN",
    "City": "Tokyo",
    "Country": "JP",
    "PostalCode": "160-0023",
    "Description": "Hilton Tokyo is located in Shinjuku, the heart of Tokyo's business, shopping and entertainment district, and is an ideal place to experience modern Japan. A complimentary shuttle operates between the hotel and Shinjuku station and the Tokyo Metro subway is connected to the hotel. Relax in one of the modern Japanese-style rooms and admire stunning city views. The hotel offers WiFi and internet access throughout all rooms and public space.",
    "Facilities": ["Pool", "WiFi ", "BusinessCenter", "DryCleaning", " Breakfast", "Bar", "BathTub"]
  }
]

*/

type SupplierA struct {
	ID            string   `json:"Id"`
	DestinationID int      `json:"DestinationId"`
	Name          string   `json:"Name"`
	Latitude      float64  `json:"Latitude"`
	Longitude     float64  `json:"Longitude"`
	Address       string   `json:"Address"`
	City          string   `json:"City"`
	Country       string   `json:"Country"`
	PostalCode    string   `json:"PostalCode"`
	Description   string   `json:"Description"`
	Facilities    []string `json:"Facilities"`
}

// custom unmarshal function to convert latitude and longitude to float64
func (s *SupplierA) UnmarshalJSON(data []byte) error {
	type Alias SupplierA
	aux := &struct {
		*Alias
		Latitude  string `json:"Latitude"`
		Longitude string `json:"Longitude"`
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		aux.Latitude = ""
		aux.Longitude = ""
	}

	// convert latitude and longitude to float64
	if aux.Latitude == "" {
		s.Latitude = 0.0
	} else {
		lat, err := strconv.ParseFloat(aux.Latitude, 64)
		if err != nil {
			s.Latitude = 0.0
		} else {
			s.Latitude = lat
		}
	}

	if aux.Longitude == "" {
		s.Longitude = 0.0
	} else {
		lat, err := strconv.ParseFloat(aux.Longitude, 64)
		if err != nil {
			s.Longitude = 0.0
		} else {
			s.Longitude = lat
		}
	}
	return nil
}

// write the adapter for supplier A
type MockySupplierA struct {
	MockySupplier
}

func (s *MockySupplierA) FetchData() ([]models.Hotel, error) {
	var hotels []models.Hotel
	var supplierA []SupplierA

	resp, err := s.fetchFromEndpoint()
	if err != nil {
		return hotels, err
	}

	// parse the response into the supplierA data format
	err = json.Unmarshal(resp, &supplierA)
	if err != nil {
		return hotels, err
	}

	// convert the supplierA data to the common data format
	for _, hotel := range supplierA {
		hotels = append(hotels, models.Hotel{
			HotelID:       hotel.ID,
			DestinationID: hotel.DestinationID,
			Name:          hotel.Name,
			Location: models.Location{
				Lat:        hotel.Latitude,
				Lng:        hotel.Longitude,
				Address:    hotel.Address,
				City:       hotel.City,
				Country:    hotel.Country,
				PostalCode: hotel.PostalCode,
			},
			Description: hotel.Description,
			Amenities: models.Amenities{
				General: hotel.Facilities,
			},
		})
	}

	return hotels, nil
}
