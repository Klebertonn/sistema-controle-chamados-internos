package mock

import "sistemas-controle-chamdos-internos/internal/models"

type MockChamadoRespository struct{}

func (m *MockChamadoRespository) Criar(Chamado models.Chamado) error {
	return nil
}

func (m *MockChamadoRespository) Atualizar(Chamado models.Chamado) error {
	return nil
}

func (m *MockChamadoRespository) Listar() ([]models.Chamado, error) {

	return []models.Chamado{
		{
			OrdemID: 1,
			Titulo:  "Test",
		},
	}, nil
}

func (m *MockChamadoRespository) BuscarOrdemID(ordemID int) (*models.Chamado, error) {

	return &models.Chamado{
		OrdemID: ordemID,
		Titulo:  "Teste",
	}, nil

}
