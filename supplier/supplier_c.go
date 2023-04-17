package supplier

import (
	"ascenda-assignment/models"
	"encoding/json"
)

/*
supplier C data: [
  {
    "id": "iJhz",
    "destination": 5432,
    "name": "Beach Villas Singapore",
    "lat": 1.264751,
    "lng": 103.824006,
    "address": "8 Sentosa Gateway, Beach Villas, 098269",
    "info": "Located at the western tip of Resorts World Sentosa, guests at the Beach Villas are guaranteed privacy while they enjoy spectacular views of glittering waters. Guests will find themselves in paradise with this series of exquisite tropical sanctuaries, making it the perfect setting for an idyllic retreat. Within each villa, guests will discover living areas and bedrooms that open out to mini gardens, private timber sundecks and verandahs elegantly framing either lush greenery or an expanse of sea. Guests are assured of a superior slumber with goose feather pillows and luxe mattresses paired with 400 thread count Egyptian cotton bed linen, tastefully paired with a full complement of luxurious in-room amenities and bathrooms boasting rain showers and free-standing tubs coupled with an exclusive array of ESPA amenities and toiletries. Guests also get to enjoy complimentary day access to the facilities at Asia’s flagship spa – the world-renowned ESPA.",
    "amenities": ["Aircon", "Tv", "Coffee machine", "Kettle", "Hair dryer", "Iron", "Tub"],
    "images": {
      "rooms": [
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg", "description": "Double room" },
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/4.jpg", "description": "Bathroom" }
      ],
      "amenities": [
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/0.jpg", "description": "RWS" },
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/6.jpg", "description": "Sentosa Gateway" }
      ]
    }
  },
  {
    "id": "f8c9",
    "destination": 1122,
    "name": "Hilton Tokyo Shinjuku",
    "lat": 35.6926,
    "lng": 139.690965,
    "address": null,
    "info": null,
    "amenities": null,
    "images": {
      "rooms": [
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i10_m.jpg", "description": "Suite" },
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i11_m.jpg", "description": "Suite - Living room" }
      ],
      "amenities": [
        { "url": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i57_m.jpg", "description": "Bar" }
      ]
    }
  }
]

*/

type SupplierCImage struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

type SupplierC struct {
	ID          string   `json:"id"`
	Destination int      `json:"destination"`
	Name        string   `json:"name"`
	Lat         float64  `json:"lat"`
	Lng         float64  `json:"lng"`
	Address     string   `json:"address"`
	Info        string   `json:"info"`
	Amenities   []string `json:"amenities"`
	Images      struct {
		Rooms     []SupplierCImage `json:"rooms"`
		Amenities []SupplierCImage `json:"amenities"`
	} `json:"images"`
}

type MockySupplierC struct {
	MockySupplier
}

func (s *MockySupplierC) FetchData() ([]models.Hotel, error) {
	var hotels []models.Hotel
	var supplierCs []SupplierC

	data, err := s.fetchFromEndpoint()
	if err != nil {
		return nil, err
	}

	// parse the data into supplierC format
	err = json.Unmarshal(data, &supplierCs)
	if err != nil {
		return nil, err
	}

	for _, d := range supplierCs {
		rooms := []models.Image{}
		amenities := []models.Image{}

		// convert the images
		for _, r := range d.Images.Rooms {
			rooms = append(rooms, models.Image{
				Link:        r.URL,
				Description: r.Description,
			})
		}
		for _, a := range d.Images.Amenities {
			amenities = append(amenities, models.Image{
				Link:        a.URL,
				Description: a.Description,
			})
		}

		hotels = append(hotels, models.Hotel{
			HotelID:       d.ID,
			DestinationID: d.Destination,
			Name:          d.Name,
			Location: models.Location{
				Lat:     d.Lat,
				Lng:     d.Lng,
				Address: d.Address,
			},
			Description: d.Info,
			Amenities: models.Amenities{
				General: d.Amenities,
			},
			Images: models.Images{
				Rooms:     rooms,
				Amenities: amenities,
			},
		})

	}

	return hotels, nil
}
