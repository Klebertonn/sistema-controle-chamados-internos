package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChamadoHendler struct{}

func newChamado() *ChamadoHendler {
	return &ChamadoHendler{}

}

func (h ChamadoHendler) CriarChamados(c gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Chamado criado com sucesso",
	})

}

func (h ChamadoHendler) ListarChamados(c gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "lista de chamados",
	})

}

func (h ChamadoHendler) BuscarOrdemID(c gin.Context) {
	ordemId := c.Param("ordemId")

	c.JSON(http.StatusOK, gin.H{
		"ordemId": ordemId,
		"message": "lista de chamados",
	})

}

func (h ChamadoHendler) AtualizarChamado(c gin.Context) {
	ordemId := c.Param("ordemId")

	c.JSON(http.StatusOK, gin.H{
		"ordemId": ordemId,
		"message": "Chamado atualizado",
	})

}
