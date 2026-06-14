package repositories

import "sistemas-controle-chamdos-internos/internal/models"

type ChamadoRepository interface {
	Criar(chamado models.Chamado)error
	Atualizar(chamado models.Chamado)error
	Listar(chamado models.Chamado)error
	BuscarOrdemID(chamado models.Chamado)error
}