package service

import (
	"sistemas-controle-chamdos-internos/internal/models"
	"sistemas-controle-chamdos-internos/internal/service/mock"
	"testing"
)

func TestBuscarOrdemID(t *testing.T) {

	t.Run("Sucesso", func(t *testing.T) {

		repo := mock.MockChamadoRespository{
			RetornarError: false,
		}

		service := NewChamadoService(&repo)

		chamado, err := service.BuscarOrdemID(1)

		if err != nil {
			t.Fatal("erro inesperado: &v", err)
		}

		if chamado.OrdemID != 1 {
			t.Errorf("esperado1, recebido %d", chamado.OrdemID)
		}

	})

	t.Run("Error", func(t *testing.T) {

		repo := mock.MockChamadoRespository{
			RetornarError: true,
		}

		service := NewChamadoService(&repo)

		chamado, err := service.BuscarOrdemID(1999)

		if err == nil {
			t.Fatal("era esperado um erro")
		}

		if chamado != nil {
			t.Fatal("chamado deveria ser nil")
		}

	})

}

func TestListarChamados(t *testing.T) {

	t.Run("Sucesso", func(t *testing.T) {

		repo := mock.MockChamadoRespository{
			RetornarError: false,
		}

		service := NewChamadoService(&repo)

		chamados, err := service.ListarChamados()

		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		if len(chamados) != 1 {
			t.Errorf("esperado 1 chamado, recebido %d", len(chamados))
		}
	})

	t.Run("Erro", func(t *testing.T) {

		repo := mock.MockChamadoRespository{
			RetornarError: true,
		}

		service := NewChamadoService(&repo)

		chamados, err := service.ListarChamados()

		if err == nil {
			t.Fatal("era esperado um erro")
		}

		if chamados != nil {
			t.Fatal("esperava lista nil")
		}
	})
}

func TestAtualizarChamado(t *testing.T) {

	t.Run("Sucesso", func(t *testing.T) {

		repo := mock.MockChamadoRespository{
			RetornarError: false,
		}

		service := NewChamadoService(&repo)

		chamado := models.Chamado{
			OrdemID:       1,
			Titulo:        "Novo teste",
			Descricao:     "TESTE ATUALIZAR",
			Status:        "Em andamento",
			Prioridade:    "Alta",
			Solicitante:   "Analista",
			ResponsavelID: 1,
		}
		err := service.AtualizarChamado(chamado)

		if err != nil {
			t.Fatalf("Error inesperado: %v", err)
		}

	})

	t.Run("Error", func(t *testing.T) {

		repo := mock.MockChamadoRespository{
			RetornarError: true,
		}

		service := NewChamadoService(&repo)

		chamado := models.Chamado{
			OrdemID: 1,
		}
		err := service.AtualizarChamado(chamado)

		if err == nil {
			t.Fatal("Error esperado")
		}

	})
}

func TestCriarChamado(t *testing.T) {

	t.Run("Sucesso", func(t *testing.T) {

		repo := mock.MockChamadoRespository{
			RetornarError: false,
		}

		service := NewChamadoService(&repo)

		chamado := models.Chamado{
			OrdemID:       1,
			Titulo:        "Criar teste",
			Descricao:     "Teste Criar test ",
			Status:        "Em andamento",
			Prioridade:    "Alta",
			Solicitante:   "Analista",
			ResponsavelID: 1,
		}
		err := service.AtualizarChamado(chamado)

		if err != nil {
			t.Fatalf("Error inesperado: %v", err)
		}

	})

	t.Run("Error", func(t *testing.T) {

		repo := mock.MockChamadoRespository{
			RetornarError: true,
		}

		service := NewChamadoService(&repo)

		chamado := models.Chamado{
			OrdemID:       1,
			Titulo:        "Criar teste",
			Descricao:     "Teste Criar test ",
			Status:        "Em andamento",
			Prioridade:    "Alta",
			Solicitante:   "Analista",
			ResponsavelID: 1,
		}

		err := service.CriarChamados(chamado)

		if err != nil {
			t.Fatal("erro esperado")
		}

	})
}
