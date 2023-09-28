package models

import (
	"db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConnectComBancoDeDados()

	selectTodosOsProdutos, err := db.Query("select * from produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		p.Id = id

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func CriarNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConnectComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(idDoProduto string) {
	db := db.ConnectComBancoDeDados()

	deletaDadosNoBanco, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deletaDadosNoBanco.Exec(idDoProduto)

	defer db.Close()
}

func EditarProduto(idDoProduto string) Produto {
	db := db.ConnectComBancoDeDados()

	produtoDadosNoBanco, err := db.Query("select * from produtos where id=$1", idDoProduto)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDadosNoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err2 := produtoDadosNoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err2 != nil {
			panic(err2.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}

	defer db.Close()

	return produtoParaAtualizar
}

func AtualizarNovoProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConnectComBancoDeDados()

	atualizarDadosNoBanco, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	atualizarDadosNoBanco.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
