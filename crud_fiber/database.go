package main

import "database/sql"

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydb"
	username     = "admin"
	password     = "admin123"
)

var db *sql.DB

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
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

func getAllProduct() ([]Product, error) {
	row, err := db.Query("SELECT id,name,price FROM product")
	if err != nil {
		return nil, err
	}
	var products []Product
	for row.Next() {
		var p Product
		err := row.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err = row.Err(); err != nil {
		return nil, err
	}
	return products, nil

}
