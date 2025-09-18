package dto

import "time"

type PerkJSON struct {
	Id        int       `json:"id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PerkType  struct {
		Id        int       `json:"id"`
		Key       string    `json:"key"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"perk_type"`
}