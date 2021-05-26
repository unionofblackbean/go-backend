package entities

import "github.com/google/uuid"

type Bot struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
