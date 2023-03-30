package inmem

import (
	"github.com/adefirmanf/travel/internal/search"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

// Search .
type Search struct{}

// SearchQuery .
func (s *Search) SearchQuery(departureCode string, departureDate string) ([]*search.Response, error) {

	q := gojsonq.New().File("list.json")
	q.Macro("date<=", dateLessOrEqualTo)
	q.Macro("date>=", dateGreaterOrEqualTo)

	var flights []*search.Response
	q.Where("departure_code", "=", departureCode).Where("departure_date", "date>=", departureDate).Out(&flights)

	return flights, nil
}

// NewSearch .
func NewSearch() *Search {
	return &Search{}
}
