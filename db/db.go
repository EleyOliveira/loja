package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBancoDados() *sql.DB {
	conexao := ""
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
