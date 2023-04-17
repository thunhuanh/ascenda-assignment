package supplier

import (
	"ascenda-assignment/models"
	"errors"
	"io"
	"net/http"
)

type Supplier interface {
	FetchData() ([]models.Hotel, error)
}

type SupplierFactory interface {
	CreateSupplier() (Supplier, error)
}

type MockySupplierFactory struct {
	Name     string
	Endpoint string
}

func (f *MockySupplierFactory) CreateSupplier() (Supplier, error) {
	switch f.Name {
	case "supplier_a":
		return &MockySupplierA{MockySupplier{Endpoint: f.Endpoint}}, nil
	case "supplier_b":
		return &MockySupplierB{MockySupplier{Endpoint: f.Endpoint}}, nil
	case "supplier_c":
		return &MockySupplierC{MockySupplier{Endpoint: f.Endpoint}}, nil
	default:
		return nil, errors.New("invalid supplier name")
	}
}

type MockySupplier struct {
	Supplier
	Endpoint string
}

func (s *MockySupplier) fetchFromEndpoint() ([]byte, error) {
	resp, err := http.Get(s.Endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error while fetching data from endpoint " + s.Endpoint)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
