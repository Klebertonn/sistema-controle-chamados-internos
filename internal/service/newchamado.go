package service

import (
	"sistemas-controle-chamdos-internos/internal/repositories"
)

func NewChamado(chamadoRepo repositories.ChamadoRepository) *ChamadoService {
	return &ChamadoService{
		chamadoRepo: chamadoRepo,
	}

}
