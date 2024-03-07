package models

import (
	"log"
	"loja/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func ConsultaProdutos() []Produto {

	db := db.ConectaBancoDados()
	consultaProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for consultaProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = consultaProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()

	return produtos
}

func Incluir(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancoDados()

	inclusao, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		log.Println("Problemas na inclusao", err)
	}

	inclusao.Exec(nome, descricao, preco, quantidade)
	defer db.Close()

}

func Deletar(id string) {
	db := db.ConectaBancoDados()

	delecao, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}

	delecao.Exec(id)
	defer db.Close()
}

