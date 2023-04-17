package workers

// worker is a function that will be executed in a goroutine and will fetch all the hotels into from all the suppliers
// and will send them to the channel that will be consumed by other workers that will clean the data and store it in the database

import (
	"ascenda-assignment/config"
	"ascenda-assignment/models"
	"ascenda-assignment/supplier"

	"log"
)

var HotelChannel chan []models.Hotel

func init() {
	HotelChannel = make(chan []models.Hotel)
}

func GetAllSupplier() []supplier.SupplierFactory {
	suppliers := []supplier.SupplierFactory{}
	// get all the suppliers from the config
	for name, endpoint := range config.GetAllSupplierURLs() {
		suppliers = append(suppliers, &supplier.MockySupplierFactory{
			Name:     name,
			Endpoint: endpoint,
		})
	}
	return suppliers
}

func UpdateData() {
	suppliers := GetAllSupplier()
	// create a goroutine for each supplier and fetch the data
	hotels := []models.Hotel{}
	dataChan := make(chan []models.Hotel)
	for _, supplierFactory := range suppliers {
		go fetchHotels(supplierFactory, dataChan)
	}

	// append the data from the channel to the hotels slice
	for i := 0; i < len(suppliers); i++ {
		hotels = append(hotels, <-dataChan...)
	}

	// send the hotels to the channel
	HotelChannel <- hotels

}

func fetchHotels(supplierFactory supplier.SupplierFactory, ch chan []models.Hotel) {
	supplier, err := supplierFactory.CreateSupplier()
	if err != nil {
		log.Println(err)
		return
	}

	hotels, err := supplier.FetchData()
	if err != nil {
		log.Println(err)
		return
	}

	// send the hotels to the channel
	ch <- hotels
}
