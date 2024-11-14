package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydb"
	username     = "admin"
	password     = "admin123"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName,
	)
	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err = sdb.Ping(); err != nil {
		log.Fatal(err)
	} else {
		println("connection success")
	}

	db = sdb

	product, err := createProduct(&Product{Name: "Smart phone", Price: 499})

	if err == nil {
		fmt.Println(product)
	}
	err = deleteProduct(product.ID)
	if err == nil {
		fmt.Println("delete success")
	}

}

func createProduct(product *Product) (Product, error) {
	var p Product
	row := db.QueryRow(
		"INSERT INTO public.product(  name, price) VALUES ($1, $2) RETURNING id,name,price;",
		product.Name, product.Price)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow("SELECT id,name,price FROM product WHERE id=$1;", id)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func updateProduct(id int, product *Product) (Product, error) {
	var p Product
	row := db.QueryRow(
		"UPDATE public.product SET  name=$1, price=$2  WHERE id=$3 RETURNING id,name,price;",
		product.Name, product.Price, id,
	)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func deleteProduct(id int) error {
	_, err := db.Exec("DELETE FROM product WHERE id=$1;", id)
	return err
}
