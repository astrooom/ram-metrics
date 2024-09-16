package main

import (
	"log"
	"net/http"

	"github.com/astrooom/ram-usage-api/internal/api"
	"github.com/astrooom/ram-usage-api/internal/config"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load()
	handler := api.NewHandler(cfg)
	r := mux.NewRouter()

	// Register RAM usage handler with query parameter
	r.HandleFunc("/", handler.RAMUsage).
		Methods("GET").
		Queries("unit", "{unit}")

	// Also allow requests without the unit query parameter
	r.HandleFunc("/", handler.RAMUsage).Methods("GET")

	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
