package config

import ("database/sql"
_"github.com/microsoft/go-mssqldb"
)
func Connect() (*sql.DB, error) {
	connString := "sqlserver://sa:!Admin123@localhost:1433?database=sistemaChamados"

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
