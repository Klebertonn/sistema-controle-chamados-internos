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

	query := `
	  INSERT INTO Chamados(
	       Titulo,
	       Descricao,
	       Status,
	       Prioridade,
	       Solicitante,
	       ResponsavelID,   
	       DataAbertura
)
	  VALUES(
	      @Titulo,
	      @Descricao,
	      @Status,
	      @Prioridade,
	      @Solicitante,
	      @ResponsavelID,   
	      GETDATE()
	)
`

	_, err := r.db.Exec(
		query,
		sql.Named("Titulo", chamado.Titulo),
		sql.Named("Descricao", chamado.Descricao),
		sql.Named("Status", chamado.Status),
		sql.Named("Prioridade", chamado.Prioridade),
		sql.Named("Solicitante", chamado.Solicitante),
		sql.Named("ResponsavelID", chamado.ResponsavelID),
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *ChamadoSQLRepository) Listar() ([]models.Chamado, error) {

	query := `
	  SELECT 
	       OrdemID,
	       Titulo,
	       Descricao,
	       Status,
	       Prioridade,
	       Solicitante,
	       ResponsavelID,   
	       DataAbertura,
	       DescCancelado
        FROM Chamados
	`

	lista, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer lista.Close()

	var chamados []models.Chamado

	for lista.Next() {
		var c models.Chamado

		err := lista.Scan(
			&c.OrdemID,
			&c.Titulo,
			&c.Descricao,
			&c.Status,
			&c.Prioridade,
			&c.ResponsavelID,
			&c.DataAbertura,
			&c.DescCancelado,
		)

		if err != nil {
			return nil, err
		}

		chamados = append(chamados, c)
	}

	return chamados, nil
}


func (r *ChamadoSQLRepository) BuscarOrdemID(ordemID int) (*models.Chamado, error) {

	query := `
	  SELECT 
	       OrdemID,
	       Titulo,
	       Descricao,
	       Status,
	       Prioridade,
	       Solicitante,
	       ResponsavelID,   
	       DataAbertura
        FROM Chamados
		WHERE OrdemID = @OrdemID
	`

	var c models.Chamado

	err := r.db.QueryRow(query, sql.Named("OrdemID", ordemID)).Scan(
		&c.OrdemID,
		&c.Titulo,
		&c.Descricao,
		&c.Status,
		&c.Prioridade,
		&c.ResponsavelID,
		&c.DataAbertura,
	)

	if err != nil {
		return nil, err
	}

	return &c, nil

}

func (r *ChamadoSQLRepository) Atualizar(chamado models.Chamado) error {

	query := `
	  UPDATE Chamados
	  SET
	      Titulo = @Titulo,
		  Descricao = @Descricao,
		  Status = @Status,
		  Prioridade = @Prioridade,
		  ResponsavelID = @ResponsavelID,
	  WHERE OrdemID = @Ordem	  	
	`

	_, err := r.db.Exec(
		query,
		sql.Named("Titulo", chamado.Titulo),
		sql.Named("Descricao", chamado.Descricao),
		sql.Named("Status", chamado.Status),
		sql.Named("Prioridade", chamado.Prioridade),
		sql.Named("Solicitante", chamado.Solicitante),
		sql.Named("ResponsavelID", chamado.ResponsavelID),
		sql.Named("OrdemID", chamado.OrdemID),
	)

	return err
}
