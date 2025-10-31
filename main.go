package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/machineash/choose-your-superhero/api/handlers"
	"github.com/machineash/choose-your-superhero/api/models"
)

var Options models.Categories

func main() {

	// load datasets at startup
	loadOptions()

	//set up routes
	http.HandleFunc("/options", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := json.MarshalIndent(Options, "", "  ")
		if err != nil {
			http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	// wiring
	http.HandleFunc("/create-hero", handlers.CreateHero)
	http.HandleFunc("/heroes", handlers.ListHeroes)

	fmt.Println("Welcome. Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func loadOptions() {
	data, err := os.ReadFile("datasets/categories.json")
	if err != nil {
		fmt.Println("error reading dataset:", err)
		return
	}
	if err := json.Unmarshal(data, &Options); err != nil {
		fmt.Println("error parsing dataset:", err)
	}
}
