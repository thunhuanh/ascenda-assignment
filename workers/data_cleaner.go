package workers

import (
	"ascenda-assignment/constants"
	"ascenda-assignment/models"
	"ascenda-assignment/repos"
	"ascenda-assignment/utils"
	"context"
	"strings"

	"fmt"
	"time"
)

func ParserAndCleaner() {
	fmt.Println("ParserAndCleaner started at: ", time.Now())

	hotelRepo := repos.NewHotelRepo(context.Background())

	for {
		// get data from hotel channel, clean it and store it in the database
		hotels := <-HotelChannel

		// merge the data from hotels that have the same ID
		// if the same field has different values, we could use machine learning model to score the text and choose the higher one
		// but for simplicity, we will just use the longest one
		hotelMapping := make(map[string][]models.Hotel)
		for _, h := range hotels {
			hotelMapping[h.HotelID] = append(hotelMapping[h.HotelID], h)
		}

		// merge the data
		mergedHotel := make([]models.Hotel, 0)
		for _, hotels := range hotelMapping {
			// get the first hotel
			hotel := hotels[0]
			for _, h := range hotels[1:] {
				hotel = Merge(hotel, h)
			}
			mergedHotel = append(mergedHotel, hotel)
		}

		// loop through the data and clean it
		for i, h := range mergedHotel {
			// clean the data
			tmpAddress := utils.StringWrapper(h.Location.Address).Clean().ToUpperAfterPeriod().ToString()
			tmpCity := utils.StringWrapper(h.Location.City).Clean().ToTitleCase().ToString()
			tmpCountry := utils.StringWrapper(h.Location.Country).Clean().ToTitleCase().ToString()

			h.Location = models.Location{
				Address: tmpAddress,
				City:    tmpCity,
				Country: tmpCountry,
			}

			h.Description = utils.StringWrapper(h.Description).Clean().ToUpperAfterPeriod().ToString()

			h.Name = utils.StringWrapper(h.Name).Clean().ToString()

			for i := range h.Amenities.General {
				h.Amenities.General[i] = utils.StringWrapper(h.Amenities.General[i]).Clean().ToString()
			}
			for i := range h.Amenities.Room {
				h.Amenities.Room[i] = utils.StringWrapper(h.Amenities.Room[i]).Clean().ToString()
			}
			for i := range h.BookingConditions {
				h.BookingConditions[i] = utils.StringWrapper(h.BookingConditions[i]).Clean().ToUpperAfterPeriod().ToString()
			}
			for i := range h.Images.Rooms {
				h.Images.Rooms[i].Description = utils.StringWrapper(h.Images.Rooms[i].Description).Clean().ToTitleCase().ToString()
			}
			for i := range h.Images.Site {
				h.Images.Site[i].Description = utils.StringWrapper(h.Images.Site[i].Description).Clean().ToTitleCase().ToString()
			}
			for i := range h.Images.Amenities {
				h.Images.Amenities[i].Description = utils.StringWrapper(h.Images.Amenities[i].Description).Clean().ToTitleCase().ToString()
			}

			mergedHotel[i] = h
		}

		// store the data in the database
		for _, h := range mergedHotel {
			err := hotelRepo.InsertHotel(h)
			if err != nil && !strings.Contains(err.Error(), constants.ErrDuplicateKey) {
				fmt.Println(err)
			}
		}
	}
}

// function to merge two hotels with the same ID
func Merge(hotel, h models.Hotel) models.Hotel {
	tmpHotel := models.Hotel{
		HotelID:       hotel.HotelID,
		DestinationID: hotel.DestinationID,
	}
	// merge the data
	tmpHotel.Name = utils.StringWrapper(hotel.Name).Merge(h.Name).ToString()
	tmpHotel.Description = utils.StringWrapper(hotel.Description).Merge(h.Description).ToString()
	tmpHotel.Location = models.Location{
		Address:    utils.StringWrapper(hotel.Location.Address).Merge(h.Location.Address).ToString(),
		City:       utils.StringWrapper(hotel.Location.City).Merge(h.Location.City).ToString(),
		Country:    utils.StringWrapper(hotel.Location.Country).Merge(h.Location.Country).ToString(),
		PostalCode: utils.StringWrapper(hotel.Location.PostalCode).Merge(h.Location.PostalCode).ToString(),
		// for lat and long, we will use whatever is available of the two
		Lat: utils.NonEmpty(h.Location.Lat, hotel.Location.Lat),
		Lng: utils.NonEmpty(h.Location.Lng, hotel.Location.Lng),
	}

	// merge the amenities
	generals := utils.RemoveDupStr(append(hotel.Amenities.General, h.Amenities.General...))
	rooms := utils.RemoveDupStr(append(hotel.Amenities.Room, h.Amenities.Room...))
	tmpHotel.Amenities = models.Amenities{
		General: generals,
		Room:    rooms,
	}

	// merge the booking conditions
	bookingConditions := utils.RemoveDupStr(append(hotel.BookingConditions, h.BookingConditions...))
	tmpHotel.BookingConditions = bookingConditions

	// merge the images
	roomImgs := make(map[string]models.Image)
	for _, img := range append(hotel.Images.Rooms, h.Images.Rooms...) {
		roomImgs[img.Link] = img
	}

	siteImgs := make(map[string]models.Image)
	for _, img := range append(hotel.Images.Site, h.Images.Site...) {
		siteImgs[img.Link] = img
	}
	amenitiesImgs := make(map[string]models.Image)
	for _, img := range append(hotel.Images.Amenities, h.Images.Amenities...) {
		amenitiesImgs[img.Link] = img
	}

	roomImgList := make([]models.Image, 0)
	for _, img := range roomImgs {
		roomImgList = append(roomImgList, img)
	}
	siteImgList := make([]models.Image, 0)
	for _, img := range siteImgs {
		siteImgList = append(siteImgList, img)
	}
	amenitiesImgList := make([]models.Image, 0)
	for _, img := range amenitiesImgs {
		amenitiesImgList = append(amenitiesImgList, img)
	}

	tmpHotel.Images = models.Images{
		Rooms:     roomImgList,
		Site:      siteImgList,
		Amenities: amenitiesImgList,
	}

	return tmpHotel
}
