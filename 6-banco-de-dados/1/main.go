package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Notebook2", 1899.99)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	productAtualizacao := NewProduct("ProdutoAtualizacao", 100.00)
	productAtualizacao.ID = "4a548b6d-6b5a-4eb1-9b17-a3318c0dd640"

	err = updateProduct(db, productAtualizacao)
	if err != nil {
		panic(err)
	}

	product.Price = 0.00
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := selectProduct(db, "4a548b6d-6b5a-4eb1-9b17-a3318c0dd640")
	if err != nil {
		panic(err)
	}
	fmt.Printf("O Produto %v, possui o valor R$ %2.f\n", p.Name, p.Price)

	products, err := selectAllProduct(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("O Produto %v, possui o valor R$ %2.f com o ID: %v\n", p.Name, p.Price, p.ID)
	}

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

}
func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price, create_at) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ?, update_at = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, time.Now(), product.ID)
	if err != nil {
		return err
	}
	return nil
}

// func selectProduct(ctx context.Context, db *sql.DB, id string) (*Product, error) {
func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	//err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil

}

func selectAllProduct(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
