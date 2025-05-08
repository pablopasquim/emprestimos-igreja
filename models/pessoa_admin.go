package models

type Pessoa struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Endereco string `json:"endereco"`
	IgrejaID       uint   `json:"igrejaID"`
}
