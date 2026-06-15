package repositories

import (
	"database/sql"
	"sistemas-controle-chamdos-internos/internal/models"
)

type ChamadoSQLRepository struct {
	db *sql.DB
}

func NewChamadoRepository(db *sql.DB) *ChamadoSQLRepository {
	return &ChamadoSQLRepository{
		db: db,
	}
}

func (r *ChamadoSQLRepository) Criar(chamado models.Chamado) error {
	return nil
}

func (r *ChamadoSQLRepository) Atualizar(chamado models.Chamado) error {
	return nil
}

func (r *ChamadoSQLRepository) Listar() ([]models.Chamado, error) {
	return nil, nil
}

func (r *ChamadoSQLRepository) BuscarOrdemID(ordemID int) (*models.Chamado, error) {
	return nil, nil
}
