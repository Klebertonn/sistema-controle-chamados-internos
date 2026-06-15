
CREATE TABLE Chamados(
    OrdemID       INT IDENTITY(4,4) PRIMARY KEY,
	Titulo        VARCHAR(200),
	Descricao     VARCHAR(MAX),
	Status        VARCHAR(50),
	Prioridade    VARCHAR(50),
	Solicitante   VARCHAR(200),
	ResponsavelID int    ,
	Data          DATETIME NOT NULL
	DescCancelado VARCHAR(MAX),


    CONSTRAINT FK_chamados_Responsaveis
    FOREIGN KEY (ResponsavelID)
    REFERENCES Responsaveis(ID)
);

CREATE INDEX IX_Chamados_Status
ON Chamados(Status);

CREATE INDEX IX_Chamados_Responsavel
ON Chamados(ResponsavelID);