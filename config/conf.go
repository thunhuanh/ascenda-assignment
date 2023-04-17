package config

var supplierURLs = map[string]string{
	"supplier_a": "http://www.mocky.io/v2/5ebbea002e000054009f3ffc",
	"supplier_b": "http://www.mocky.io/v2/5ebbea102e000029009f3fff",
	"supplier_c": "http://www.mocky.io/v2/5ebbea1f2e00002b009f4000",
}

const DB_URI = "mongodb+srv://test:0ay38qfqzF20djrI@test.mtfiuuf.mongodb.net/?retryWrites=true&w=majority"

func GetSupplierURL(name string) string {
	return supplierURLs[name]
}

func GetAllSupplierURLs() map[string]string {
	return supplierURLs
}
