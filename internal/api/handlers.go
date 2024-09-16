package api

import (
	"encoding/json"
	"net/http"

	"github.com/astrooom/ram-usage-api/internal/config"
	"github.com/astrooom/ram-usage-api/internal/ram"
	"github.com/astrooom/ram-usage-api/pkg/auth"
	"github.com/gorilla/mux"
)

type Handler struct {
	config *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{config: cfg}
}

func (h *Handler) RAMUsage(w http.ResponseWriter, r *http.Request) {
	if !auth.ValidateToken(r.Header.Get("Authorization"), h.config.APIToken) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	usage, err := ram.GetUsage()
	if err != nil {
		http.Error(w, "Failed to get RAM usage", http.StatusInternalServerError)
		return
	}

	// Get the unit from query parameters
	vars := mux.Vars(r)
	unit := vars["unit"]

	// Convert usage based on the unit
	convertedUsage := convertUsage(usage, unit)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(convertedUsage)
}

func convertUsage(usage *ram.Usage, unit string) *ram.Usage {
	var divisor uint64
	switch unit {
	case "kb":
		divisor = 1024
	case "mb":
		divisor = 1024 * 1024
	case "gb":
		divisor = 1024 * 1024 * 1024
	default:
		return usage // Return original usage if unit is not specified or invalid
	}

	if divisor > 1 {
		usage.Total /= divisor
		usage.Available /= divisor
		usage.Used /= divisor
		usage.Free /= divisor
	}

	return usage
}
