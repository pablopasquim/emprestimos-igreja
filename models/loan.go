package models

type Loan struct {
	ID             int    `json:"id"`
	ItemID         uint   `json:"item_id"`
	FavoredID      uint   `json:"pessoa_id"`
	ChurchID       uint   `json:"church"`
	DataLoan       string `json:"data_loan"`
	DataDevolution string `json:"data_devolution"`
}
