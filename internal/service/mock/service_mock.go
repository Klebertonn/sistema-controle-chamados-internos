package mock

import (
	"errors"
	"sistemas-controle-chamdos-internos/internal/models"
)

type MockChamadoRespository struct {
	RetornarError bool
}

func (m *MockChamadoRespository) Criar(chamado models.Chamado) error {
	return nil
}

func (m *MockChamadoRespository) Atualizar(Chamado models.Chamado) error {
	return nil
}

func (m *MockChamadoRespository) Listar() ([]models.Chamado, error) {

	return []models.Chamado{
		{
			OrdemID:       1,
			Titulo:        "Test",
			Descricao:     "Test service",
			Status:        "Aberto",
			Prioridade:    "Alta",
			Solicitante:   "Analista",
			ResponsavelID: 1,
		},
	}, nil
}

func (m *MockChamadoRespository) BuscarOrdemID(ordemID int) (*models.Chamado, error) {

	if m.RetornarError {
		return nil, errors.New("Chamado não encontrado")
	}

	return &models.Chamado{
		OrdemID: ordemID,
		Titulo:  "Teste",
	}, nil

}
