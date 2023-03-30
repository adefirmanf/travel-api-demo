package inmem

import (
	"log"
	"testing"
)

func TestSearchQuery(t *testing.T) {
	s := NewSearch()
	data, err := s.SearchQuery("CGK", "2023-03-30")
	if err != nil {
		t.Fail()
	}

	for _, v := range data {
		log.Println(v.Price)
	}
}
