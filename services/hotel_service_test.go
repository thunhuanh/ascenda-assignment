package services

import (
	"ascenda-assignment/models"
	"ascenda-assignment/repos"
	"context"
	"log"
	"testing"
)

var hotelRepo repos.HotelRepo

func TestMain(m *testing.M) {
	// setup repo
	hotelRepo = repos.NewHotelRepo(context.Background())
	m.Run()
}
func Test_hotelService_FindHotel(t *testing.T) {

	type args struct {
		destination int
		hotels      []string
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.Hotel
		wantErr bool
	}{
		// test case 1 - valid destination
		{
			name: "valid destination",
			args: args{
				destination: 5432,
			},
			want:    make([]*models.Hotel, 2), // 2 hotels in the database have destination 5432
			wantErr: false,
		},
		// test case 2 - pass in invalid destination and hotels
		{
			name: "hotels not found",
			args: args{
				destination: 1,
				hotels:      []string{"hotel1", "hotel2", "hotel3"},
			},
			want:    []*models.Hotel{},
			wantErr: true,
		},
		// test case 3 - just one hotel found
		{
			name: "just one hotel found",
			args: args{
				hotels: []string{"iJhz"},
			},
			want:    make([]*models.Hotel, 1),
			wantErr: false,
		},
		// test case 4 - pass in valid hotels only
		{
			name: "pass in valid hotels only",
			args: args{
				hotels: []string{"iJhz", "SjyX"},
			},
			want:    make([]*models.Hotel, 2),
			wantErr: false,
		},
		// test case 5 - pass in valid destination and invalid hotels
		{
			name: "pass in valid destination and invalid hotels",
			args: args{
				destination: 5432,
				hotels:      []string{"invalid1", "invalid2", "invalid3"},
			},
			want:    make([]*models.Hotel, 2),
			wantErr: false,
		},
		// test case 6 - pass in invalid destination and valid hotels
		{
			name: "pass in invalid destination and valid hotels",
			args: args{
				destination: 1,
				hotels:      []string{"iJhz", "SjyX"},
			},
			want:    make([]*models.Hotel, 2),
			wantErr: false,
		},
		// test case 7 - pass in invalid destination only
		{
			name: "pass in invalid destination only",
			args: args{
				destination: 1,
			},
			want:    []*models.Hotel{},
			wantErr: true,
		},
		// test case 8 - pass in invalid invalid hotels only
		{
			name: "pass in invalid invalid hotels only",
			args: args{
				hotels: []string{"invalid1", "invalid2", "invalid3"},
			},
			want:    []*models.Hotel{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hotelService{
				repo: hotelRepo,
			}
			got, err := h.FindHotel(tt.args.destination, tt.args.hotels)
			if (err != nil) != tt.wantErr {
				log.Println(tt.args.destination, tt.args.hotels)
				t.Errorf("hotelService.FindHotel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !isEqual(got, tt.want) {
				t.Errorf("hotelService.FindHotel() = %v, want %v", got, tt.want)
			}
		})
	}
}

// for simplicity, we can use this function to compare two slices
func isEqual(got, want []*models.Hotel) bool {
	return len(got) == len(want)
}
