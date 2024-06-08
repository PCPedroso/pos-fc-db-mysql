package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Produto struct {
	ID    string
	Nome  string
	Preco float64
}

func NewProduct(nome string, preco float64) *Produto {
	return &Produto{
		ID:    uuid.New().String(),
		Nome:  nome,
		Preco: preco,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	produto := NewProduct("Monitor", 1050.00)
	err = inserirProduto(db, produto)

	if err != nil {
		panic(err)
	}

	//err = updateProduto(db, produto)
	//produto, err = selectProduto(db, "fb4c678e-4191-440c-8d76-192f9a26a0a4")
	//produtos, err := selectProdutos(db)

	// for _, produto := range produtos {
	// 	fmt.Printf("Produto: %v possui o valor de %.2f\n", produto.Nome, produto.Preco)
	// }

	err = deleteProduto(db, "fb4c678e-4191-440c-8d76-192f9a26a0a4")
	if err != nil {
		panic(err)
	}

}

func inserirProduto(db *sql.DB, produto *Produto) error {
	// Utilizando Prepare para evitar SQL Injection
	stmt, err := db.Prepare("insert into produto (id, nome, preco) values (?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(produto.ID, produto.Nome, produto.Preco)
	if err != nil {
		panic(err)
	}
	return nil
}

func updateProduto(db *sql.DB, produto *Produto) error {
	// Utilizando Prepare para evitar SQL Injection
	stmt, err := db.Prepare("update produto set preco = ? where nome = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(produto.Preco, produto.Nome)
	if err != nil {
		panic(err)
	}
	return nil
}

func selectProduto(db *sql.DB, id string) (*Produto, error) {
	// Utilizando Prepare para evitar SQL Injection
	stmt, err := db.Prepare("select id, nome, preco from produto where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var produto Produto

	err = stmt.QueryRow(id).Scan(&produto.ID, &produto.Nome, &produto.Preco)
	if err != nil {
		panic(err)
	}

	return &produto, nil
}

func selectProdutos(db *sql.DB) ([]Produto, error) {
	rows, err := db.Query("select id, nome, preco from produto")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var listaProdutos []Produto

	for rows.Next() {
		var produto Produto

		err = rows.Scan(&produto.ID, &produto.Nome, &produto.Preco)
		if err != nil {
			panic(err)
		}

		listaProdutos = append(listaProdutos, produto)
	}

	return listaProdutos, err
}

func deleteProduto(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from produto where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		panic(err)
	}
	return nil
}
