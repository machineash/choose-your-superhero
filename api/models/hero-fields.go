package models

import "time"

// define the data structure and fields

// define request and response

type Categories struct {
	Powers       []string `json:"powers"`
	Environments []string `json:"environments"`
	Vices        []string `json:"vices"`
	Stabilizers  []string `json:"stabilizers"`
	Stressors    []string `json:"stressors"`
	Personality  []string `json:"personality"`
}

type HeroProfile struct {
	ID          string    `'json:"id"`
	Power       string    `'json:"power"`
	Environment string    `'json:"environment"`
	Vice        string    `'json:"vice"`
	Stabilizers []string  `'json:"stabilizers"`
	Stressors   []string  `'json:"stressors"`
	Personality string    `'json:"personality"`
	ClosestHero string    `'json:"closest_hero"`
	Message     string    `'json:"message"`
	CreatedAt   time.Time `'json:"created_at"`
}

type HeroRequest struct {
	Power       string   `json:"power"`
	Environment string   `json:"environment"`
	Vice        string   `json:"vice"`
	Stabilizers []string `json:"stabilizers"`
	Stressors   []string `json:"stressors"`
	Personality string   `json:"personality"`
}

type HeroResponse struct {
	ClosestHero string `json:"closest_hero"`
	Message     string `json:"message"`
}
