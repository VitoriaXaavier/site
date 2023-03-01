package models

import "github.com/VitoriaXaavier/site/db"

type Produto struct {
	Id         int
	Nome       string
	Descrição  string
	Preço      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db := db.ConectaBancoDados()
	defer db.Close()
	selectDeTodosProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())

		}
		p.Nome = nome
		p.Id = id
		p.Descrição = descricao
		p.Preço = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancoDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}
func DeletaProduto(id string) {
	db := db.ConectaBancoDados()
	defer db.Close()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)
}

func EditaProduto (id string) Produto {
	db := db.ConectaBancoDados()
	defer db.Close()

	produtoBanco, err := db.Query("select * from produtos where id=$1",id)
	
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoBanco.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id,&nome,&descricao,&preco,&quantidade)

		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descrição = descricao
		produtoParaAtualizar.Preço = preco
		produtoParaAtualizar.Quantidade = quantidade

	}
	return produtoParaAtualizar

}

func AtualizaProduto (nome string,  descricao string, preco float64, quantidade int, id int64) {
	db := db.ConectaBancoDados()
	defer db.Close()

	atualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(nome,descricao,preco,quantidade,id)
}