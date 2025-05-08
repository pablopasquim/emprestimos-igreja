package models

type Church struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	localization string `json:"localization"`
}
