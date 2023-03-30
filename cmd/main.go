package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adefirmanf/travel/internal/search"
	"github.com/adefirmanf/travel/internal/search/inmem"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	l "github.com/ahmetb/go-linq/v3"
)

func main() {
	app := search.NewSearch(inmem.NewSearch())
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	// search?detination - required | departure_date - optional
	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		departureCode := r.URL.Query().Get("departure_code")
		departureDate := r.URL.Query().Get("departure_date")
		sortByPriceQ := r.URL.Query().Get("sort_by_price")

		resp, err := app.SearchQuery(departureCode, departureDate)
		if err != nil {
			log.Println(err)
		}

		responseWithJSON(w, http.StatusOK, sortByPrice(sortByPriceQ, resp))
	})
	http.ListenAndServe(":3000", r)
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func sortByPrice(order string, results []*search.Response) []*search.Response {
	def := "asc"
	if order == "" {
		order = def
	}
	switch order {
	case "asc":
		l.From(results).Sort(func(i interface{}, j interface{}) bool {
			return i.(*search.Response).Price < j.(*search.Response).Price
		}).ToSlice(&results)
	case "desc":
		l.From(results).Sort(func(i interface{}, j interface{}) bool {
			return i.(*search.Response).Price > j.(*search.Response).Price
		}).ToSlice(&results)
	}
	return results
}
