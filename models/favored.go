package models

type Favored struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Andress  string `json:"andress"`
	ChurchID uint   `json:"church_id"`
}
