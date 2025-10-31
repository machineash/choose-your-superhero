package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/machineash/choose-your-superhero/api/models"
	"github.com/machineash/choose-your-superhero/services"
)

func CreateHero(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("POST JSON here to create a hero."))
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.HeroRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		if errors.Is(err, io.EOF) {
			http.Error(w, "empty body: send JSON data", http.StatusBadRequest)
			return
		}
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// for Phase 1, we keep match simple
	closest := "unknown hero"
	message := "Phase 1 complete - received your hero data."

	profile := models.HeroProfile{
		Power:       req.Power,
		Environment: req.Environment,
		Vice:        req.Vice,
		Stabilizers: req.Stabilizers,
		Stressors:   req.Stressors,
		Personality: req.Personality,
		ClosestHero: closest,
		Message:     message,
	}

	saved, err := services.SaveHero(profile)
	if err != nil {
		http.Error(w, "failed to save profile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(saved)
}

func ListHeroes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	heroes, err := services.LoadHeroes()
	if err != nil {
		http.Error(w, "failed to load heroes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}
