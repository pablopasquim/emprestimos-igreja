package models

type Item struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Tipo   string `json:"tipo"`
	Status string `json:"status"`
}

type Pessoa struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Endereco string `json:"endereco"`
}

type Igreja struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	Localizacao string `json:"localizacao"`
}

type Emprestimo struct {
	ID             int    `json:"id"`
	ItemID         uint   `json:"itemID"`
	PessoaID       uint   `json:"pessoaID"`
	IgrejaID       uint   `json:"igrejaID"`
	DataEmprestimo string `json:"dataEmprestimo"`
	DataDevolucao  string `json:"dataDevolucao"`
}
