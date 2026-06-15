package service

import (
	"errors"
	"sistemas-controle-chamdos-internos/internal/models"
	"sistemas-controle-chamdos-internos/internal/repositories"
	"time"
)

type ChamadoService struct {
	chamadoRepo     repositories.ChamadoRepository
	responsavelRepo repositories.ResponvaelRepository
}

func NewChamadoService(chamadoRepo repositories.ChamadoRepository) *ChamadoService {
	return &ChamadoService{
		chamadoRepo: chamadoRepo,
	}

}

func (s ChamadoService) CriarChamados(chamado models.Chamado) error {

	if chamado.Titulo == "" {
		return errors.New("título obrigatório")
	}

	if chamado.Solicitante == "" {
		return errors.New("O solicitante é obrigatório")
	}

	chamado.Status = models.StatusAberto
	chamado.DataAbertura = time.Now()

	if chamado.ResponsavelID == 0 {
		responsavel, err := s.responsavelRepo.BuscarResp()

		if err != nil {
			return err
		}

		chamado.ResponsavelID = responsavel.ID
	}

	return s.chamadoRepo.Criar(chamado)

}

func (s ChamadoService) BuscarOrdemID(ordemID int) (*models.Chamado, error) {

	return s.chamadoRepo.BuscarOrdemID(ordemID)
}

func (s ChamadoService) ListarChamados() ([]models.Chamado, error) {

	return s.chamadoRepo.Listar()

}

func (s ChamadoService) AtualizarChamado(chamado models.Chamado) error {

	existe, err := s.chamadoRepo.BuscarOrdemID(chamado.OrdemID)

	if err != nil {
		return err
	}

	if chamado.Status == models.StatusCancelado  {
		return errors.New("Descrição do cancelamento é obrigatório ")
	}

	if chamado.ResponsavelID > 0 && chamado.Status == models.StatusAberto {
		chamado.Status = models.StatusaAndamento
	}

	chamado.DataAbertura = existe.DataAbertura

	return s.chamadoRepo.Atualizar(chamado)
}
