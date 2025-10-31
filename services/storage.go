package services

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/machineash/choose-your-superhero/api/models"
)

var (
	filePath = "artifacts/heroes.json"
	mu       sync.Mutex
)

// ensure artifacts directory exists
func ensureDir() error {
	dir := filepath.Dir(filePath)
	return os.MkdirAll(dir, 0755)
}

func LoadHeroes() ([]models.HeroProfile, error) {
	mu.Lock()
	defer mu.Unlock()

	if err := ensureDir(); err != nil {
		return nil, err
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.HeroProfile{}, nil
		}
		return nil, err
	}

	var heroes []models.HeroProfile
	_ = json.Unmarshal(b, &heroes)
	return heroes, nil
}

func SaveHero(h models.HeroProfile) (models.HeroProfile, error) {
	mu.Lock()
	defer mu.Unlock()

	if err := ensureDir(); err != nil {
		return h, err
	}

	// read file directly
	var heroes []models.HeroProfile
	if b, err := os.ReadFile(filePath); err == nil {
		_ = json.Unmarshal(b, &heroes)
	}

	// assign ID + timestamp
	h.ID = time.Now().UTC().Format("20060102T150405.000000000")
	h.CreatedAt = time.Now().UTC()

	heroes = append(heroes, h)

	b, err := json.MarshalIndent(heroes, "", "  ")
	if err != nil {
		return h, err
	}

	if err := os.WriteFile(filePath, b, 0644); err != nil {
		return h, err
	}
	return h, nil
}
