package models

type Filter struct {
	DestinationID int      `json:"destination_id,omitempty"`
	Hotels        []string `json:"hotels,omitempty"`
}
