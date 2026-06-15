package service

import (
	"sistemas-controle-chamdos-internos/internal/service/mock"
	"testing"
)

func TesteBuscarOrdemID(t *testing.T) {

	repo := mock.MockChamadoRespository{}

	service := NewChamadoService(&repo)

	chamado, err := service.BuscarOrdemID(1)

	if err != nil {
		t.Fatal("erro inesperado: &v", err)
	}

	if chamado.OrdemID != 1 {
		t.Errorf("esperado1, recebido %d", chamado.OrdemID)
	}

}
