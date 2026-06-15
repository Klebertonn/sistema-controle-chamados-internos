CREATE TABLE Chamados(
    OrdemID       INT IDENTITY(4,4) PRIMARY KEY,
	Titulo        VARCHAR(200),
	Descricao     VARCHAR(MAX),
	Status        VARCHAR(50),
	Prioridade    VARCHAR(50),
	Solicitante   VARCHAR(200),
	ResponsavelID int    
	Data          time.Time
	DescCancelado VARCHAR(MAX),


    CONSTRANT FK_chamados_Responsaveis
    FORIGN KEY (ResponsavelID)
    REFERENCES Responsaveis(ID)
);

CREATE INDEX IX_Chamados_Status
ON Chamados(Status);

CREATE INDEX IX_Chamados_Responsavel
ON Chamados(ResponsavelID);