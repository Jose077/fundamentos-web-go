package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// conexão com o banco de dados
func ConectToDb() *sql.DB {
	conexao := "user=postgres dbname=fundamento_web password=postgres host=localhost sslmode=disable port=5433"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
