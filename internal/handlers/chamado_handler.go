package handlers

import (
	"net/http"
	"sistemas-controle-chamdos-internos/internal/models"
	"sistemas-controle-chamdos-internos/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChamadoHendler struct {
	service *service.ChamadoService
}

func NewChamadoHandler(service *service.ChamadoService) *ChamadoHendler {
	return &ChamadoHendler{
		service: service,
	}

}

func (h *ChamadoHendler) CriarChamados(c *gin.Context) {

	var chamado models.Chamado

	if err := c.ShouldBindJSON(&chamado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "dados inválidos",
		})
		return
	}

	if err := h.service.CriarChamados(chamado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Chamado criado com sucesso",
	})
}

func (h *ChamadoHendler) ListarChamados(c *gin.Context) {

	chamados, err := h.service.ListarChamados()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, chamados)

}

func (h *ChamadoHendler) BuscarOrdemID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Numero de ordem inválida",
		})
		return
	}

	chamado, err := h.service.BuscarOrdemID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "chamado não econtrado",
		})
		return
	}

	c.JSON(http.StatusOK, chamado)

}

func (h *ChamadoHendler) AtualizarChamado(c *gin.Context) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Numero de ordem inválida",
		})
		return
	}

	var chamado models.Chamado

	if err := c.ShouldBindJSON(&chamado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "dados inválidos",
		})
		return
	}

	chamado.OrdemID = id

	if err := h.service.AtualizarChamado(chamado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "chamado atualizaod com sucesso",
	})

}
