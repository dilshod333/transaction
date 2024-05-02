package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	ProductID   int
	ProductName string
	CategoryID  int
	Unit        string
	Price       float64
}
const (
	host="localhost"
	user="postgres"
	password="Dilshod@2005"
	port=5432
	dbname="demo"
)
var db *sql.DB
func main() {

	var err error
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable",host, user,password,port, dbname)
	db,err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}

	new := Product{
		ProductName: "Sabzi",
		CategoryID:  97,
		Unit:        "kg",
		Price:       15.77,
	}
	_, err = tx.Exec(`INSERT INTO products (product_name, category_id, unit, price) 
						   VALUES ($1, $2, $3, $4)`, new.ProductName, new.CategoryID, new.Unit, new.Price)
	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
	}


	var productGet Product
	err = tx.QueryRow(`select * from products where product_name = $1`, new.ProductName).Scan(
		&productGet.ProductID,
		&productGet.ProductName,
		&productGet.CategoryID,
		&productGet.Unit,
		&productGet.Price,
	)
	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
	}
	fmt.Printf("get back product: %v\n", productGet)


	_, err = tx.Exec("UPDATE products SET price = $1 WHERE product_name = $2", 25.99, new.ProductName)
	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
	}


	var modifProduct Product
	err = tx.QueryRow(`select * from products where product_name = $1`, new.ProductName).Scan(
		&modifProduct.ProductID,
		&modifProduct.ProductName,
		&modifProduct.CategoryID,
		&modifProduct.Unit,
		&modifProduct.Price,
	)
	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
	}
	fmt.Printf("updated product: %v\n", modifProduct)


	_, err = tx.Exec("delete from products where product_name = $1", new.ProductName)
	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
	}


	err = tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}
