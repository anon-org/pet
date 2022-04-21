package main

import (
	"encoding/json"
	"fmt"
	"github.com/anon-org/arithmetic/api"
	"github.com/anon-org/arithmetic/impl"
	"log"
	"net/http"
)

func main() {
	svc := impl.Wire(make(map[string]*api.User))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		switch r.Method {
		case http.MethodGet:
			fetch(svc, w, r)
		case http.MethodPost:
			store(svc, w, r)
		}
	})

	fmt.Println("listening at :8000")
	http.ListenAndServe(":8000", nil)
}

func fetch(svc api.Service, w http.ResponseWriter, r *http.Request) {
	log.Println("fetching...")
	results, err := svc.Fetch(r.Context())
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(results)
	}
}

func store(svc api.Service, w http.ResponseWriter, r *http.Request) {
	var req api.Request
	json.NewDecoder(r.Body).Decode(&req)

	log.Printf("storing %s\n", req.Name)
	results, err := svc.Store(r.Context(), req.Name)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(results)
	}
}
