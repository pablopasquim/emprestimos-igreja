package models

type Loan struct {
	ID             int    `json:"id"`
	ItemID         uint   `json:"itemID"`
	PessoaID       uint   `json:"pessoaID"`
	IgrejaID       uint   `json:"igrejaID"`
	DataEmprestimo string `json:"dataEmprestimo"`
	DataDevolucao  string `json:"dataDevolucao"`
}
