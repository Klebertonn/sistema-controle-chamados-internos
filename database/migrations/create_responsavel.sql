CREATE TABLE Responsaveis{
    ID int IDENTITY(1,1) PRIMARY KEY,
    Nomes VARCHAR(100) NOT NULL
}

iNSERT INTO Responsaveis(Nomes)
VALUES
("Lucas"),
("Pedro"),
("Maria"),
("julio");