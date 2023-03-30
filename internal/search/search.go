package search

import "encoding/json"

// Service .
type Service interface {
	SearchQuery(departureCode string, departureDate string) ([]*Response, error)
}

// Search .
type Search struct {
	service Service
}

// Response .
type Response struct {
	AirlineCode          string      `json:"airline_code"`
	AirlineName          string      `json:"airline_name"`
	DepartureCode        string      `json:"departure_code"`
	DepartureName        string      `json:"departure_name"`
	ArrivalCode          string      `json:"arrival_code"`
	ArrivalName          string      `json:"arrival_name"`
	ArrivalDateEta       string      `json:"arrival_date_eta"`
	DurationEtaInMinutes string      `json:"duration_eta_in_minutes"`
	Price                json.Number `json:"price"` // Todo: should not pass the JSON.Number as data type
}

// SearchQuery .
func (s *Search) SearchQuery(departureCode string, departureDate string) ([]*Response, error) {
	return s.service.SearchQuery(departureCode, departureDate)
}

// NewSearch .
func NewSearch(s Service) *Search {
	return &Search{
		service: s,
	}
}
