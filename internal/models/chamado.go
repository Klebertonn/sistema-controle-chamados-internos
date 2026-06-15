package models

import (
	"time"
)

type Chamado struct {
	OrdemID       int       `json:"ordemId"`
	Titulo        string    `json:"titulo"`
	Descricao     string    `json:"descricao"`
	Status        string    `json:"status"`
	Prioridade    string    `json:"prioridade"`
	Solicitante   string    `json:"solicitante"`
	ResponsavelID int       `json:"responsavelId"`
	DataAbertura  time.Time `json:"dataAbertura"`
}
