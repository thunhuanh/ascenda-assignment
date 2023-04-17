package models

import (
	"ascenda-assignment/infra"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Hotel struct {
	HotelID           string    `json:"id,omitempty" bson:"hotel_id,omitempty"`
	DestinationID     int       `json:"destination_id,omitempty" bson:"destination_id,omitempty"`
	Name              string    `json:"name,omitempty" bson:"name,omitempty"`
	Location          Location  `json:"location,omitempty" bson:"location,omitempty"`
	Description       string    `json:"description,omitempty" bson:"description,omitempty"`
	Amenities         Amenities `json:"amenities,omitempty" bson:"amenities,omitempty"`
	Images            Images    `json:"images,omitempty" bson:"images,omitempty"`
	BookingConditions []string  `json:"booking_conditions,omitempty" bson:"booking_conditions,omitempty"`
}

type Location struct {
	Lat        float64 `json:"lat,omitempty" bson:"lat,omitempty"`
	Lng        float64 `json:"lng,omitempty" bson:"lng,omitempty"`
	Address    string  `json:"address,omitempty" bson:"address,omitempty"`
	City       string  `json:"city,omitempty" bson:"city,omitempty"`
	Country    string  `json:"country,omitempty" bson:"country,omitempty"`
	PostalCode string  `json:"postal_code,omitempty" bson:"postal_code,omitempty"`
}

type Amenities struct {
	General []string `json:"general,omitempty" bson:"general,omitempty"`
	Room    []string `json:"room,omitempty" bson:"room,omitempty"`
}

type Images struct {
	Rooms     []Image `json:"rooms,omitempty" bson:"rooms,omitempty"`
	Site      []Image `json:"site,omitempty" bson:"site,omitempty"`
	Amenities []Image `json:"amenities,omitempty" bson:"amenities,omitempty"`
}

type Image struct {
	Link        string `json:"link,omitempty" bson:"link,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}

func init() {
	// create indexes
	dbClient := infra.GetDBClient()
	database := dbClient.Database("ascenda")
	col := database.Collection("hotels")
	indexes := []mongo.IndexModel{
		{
			Keys: bson.M{
				"hotel_id": 1,
			},
			Options: options.Index().SetUnique(true).SetBackground(true),
		},
		{
			Keys: bson.D{
				{Key: "destination_id", Value: 1},
				{Key: "hotel_id", Value: 1},
			},
			Options: options.Index().SetBackground(true),
		},
	}
	_, err := col.Indexes().CreateMany(context.Background(), indexes)
	if err != nil {
		panic(err)
	}

}
