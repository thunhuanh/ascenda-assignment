package repos

import (
	"ascenda-assignment/infra"
	"ascenda-assignment/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

type HotelRepo interface {
	GetHotel(destination int, hotels []string) ([]*models.Hotel, error)
	InsertHotel(hotel models.Hotel) error
	InsertManyHotel(hotels []*models.Hotel) error
	DeleteHotel(hotelID string) error
}

// import gorm package
type hotelRepo struct {
	HotelRepo
	ctx context.Context
}

func (h *hotelRepo) GetHotel(destination int, hotels []string) ([]*models.Hotel, error) {
	dbClient := infra.GetDBClient()
	database := dbClient.Database("ascenda")
	col := database.Collection("hotels")

	orQuery := []bson.M{}
	destinationQuery := bson.M{}
	if destination != 0 {
		destinationQuery["destination_id"] = destination
		orQuery = append(orQuery, destinationQuery)
	}

	hotelIdQuery := bson.M{}
	if len(hotels) != 0 {
		hotelIdQuery["hotel_id"] = bson.M{
			"$in": hotels,
		}
		orQuery = append(orQuery, hotelIdQuery)
	}

	filter := bson.M{
		"$or": orQuery,
	}
	cur, err := col.Find(h.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(h.ctx)
	results := []*models.Hotel{}
	for cur.Next(h.ctx) {
		var result models.Hotel
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	if len(results) == 0 {
		return nil, errors.New("hotel not found")
	}

	return results, nil

}

func (h *hotelRepo) InsertHotel(hotel models.Hotel) error {
	dbClient := infra.GetDBClient()
	database := dbClient.Database("ascenda")
	col := database.Collection("hotels")

	_, err := col.InsertOne(h.ctx, hotel)
	if err != nil {
		return err
	}
	return nil
}

func (h *hotelRepo) InsertManyHotel(hotels []*models.Hotel) error {
	dbClient := infra.GetDBClient()
	database := dbClient.Database("ascenda")
	col := database.Collection("hotels")

	data := []interface{}{}
	for _, hotel := range hotels {
		data = append(data, hotel)
	}

	_, err := col.InsertMany(h.ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (h *hotelRepo) DeleteHotel(hotelID string) error {
	dbClient := infra.GetDBClient()
	database := dbClient.Database("ascenda")
	col := database.Collection("hotels")

	filter := bson.M{
		"hotel_id": hotelID,
	}
	_, err := col.DeleteOne(h.ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func NewHotelRepo(ctx context.Context) *hotelRepo {
	return &hotelRepo{
		ctx: ctx,
	}
}
