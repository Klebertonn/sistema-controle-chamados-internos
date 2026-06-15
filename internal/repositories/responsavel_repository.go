package repositories

import "sistemas-controle-chamdos-internos/internal/models"

type ResponvaelRepository interface {
	Listar() ([]models.Responsavel, error)
	BuscarID(id int) (*models.Responsavel, error)
	BuscarResp() (*models.Responsavel, error)
}
