package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dylanlott/meroxa-project/datasource"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("starting meroxa monitor")
	var directory = make(map[uuid.UUID]datasource.Datasource)
	http.HandleFunc("/datasource/create", func(w http.ResponseWriter, r *http.Request) {
		// grab payload from request
		s := &datasource.Source{}
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// create new *Source
		fmt.Printf("decoded source: %+v\n", s)
		id, err := uuid.NewUUID()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// persist to our directory
		s.ID = id
		directory[s.ID] = s

		// return results
	})

	http.ListenAndServe(":7070", nil)
}
