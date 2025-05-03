package models

type Item struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Tipo   string `json:"tipo"`
	Status string `json:"status"`
}
