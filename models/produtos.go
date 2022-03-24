package models

import "main/db"

type Produto struct {
	id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {

	db := db.ConectToDb()

	produtosDb, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for produtosDb.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtosDb.Scan(&id, &nome, &descricao, &preco, &quantidade)

		println("id: ", id, "nome: ", nome, "descricao: ", descricao)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos

}
