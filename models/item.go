package models

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	//values: available & unavailable
	Status string `json:"status"`
}
