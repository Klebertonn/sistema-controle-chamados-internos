package migrations

import (
	"database/sql"
	"fmt"
)

func Migrate(db *sql.DB) error {
	queryResponsaveis := `
	IF NOT EXISTS(
		SELECT * 
		FROM sys.tables
		WHERE name = 'Responsaveis'
	)
	BEGIN
	CREATE TABLE Responsaveis(
    ID int IDENTITY(1,1) PRIMARY KEY,
    Nome VARCHAR(100) NOT NULL
    )
END
`

	_, err := db.Exec(queryResponsaveis)
	if err != nil {
		return fmt.Errorf(
			"Error ao criar tabela Responsaveis:%w", err)
	}

	queryChamados := `
	IF NOT EXISTS(
		SELECT * 
		FROM sys.tables
		WHERE name = 'Chamados'
	)
	BEGIN
	  CREATE TABLE Chamados(
    OrdemID       INT IDENTITY(4,4) PRIMARY KEY,
	Titulo        VARCHAR(200) NOT NULL,
	Descricao     VARCHAR(MAX) NOT NULL,
	Status        VARCHAR(50)  NOT NULL,
	Prioridade    VARCHAR(50)  NOT NULL,
	Solicitante   VARCHAR(200) NOT NULL,
	ResponsavelID int          NULL   
	Data          DATETIME     NOT NULL
	DescCancelado VARCHAR(MAX) NULL,


    CONSTRANT FK_chamados_Responsaveis
    FORIGN KEY (ResponsavelID)
    REFERENCES Responsaveis(ID)
)
END
`

	_, err = db.Exec(queryChamados)
	if err != nil {
		return fmt.Errorf(
			"Error ao criar tabela Chamados:%w", err)
	}

	queryIndexStatus := `
	IF NOT EXISTS(
		SELECT * 
		FROM sys.indexes
		WHERE name = 'IX_chamados_Status'
	)
	CREATE INDEX IX_Chamados_Status
	ON Chamados(Status)
    `

	_, err = db.Exec(queryIndexStatus)
	if err != nil {
		return fmt.Errorf(
			"Error ao gerar o indice Status: %w", err)
	}

	queryIndexResponsavel := `
	IF NOT EXISTS(
		SELECT * 
		FROM sys.indexes
		WHERE name = 'IX_chamados_Responsavel'
	)
	CREATE INDEX IX_Chamados_Responsavel
	ON Chamados(ResponsavelID)
    `

	_, err = db.Exec(queryIndexResponsavel)
	if err != nil {
		return fmt.Errorf(
			"Error ao gerar o indice Status: %w", err)
	}

	fmt.Println(" Migração executada com sucesso!")
	return nil

}
