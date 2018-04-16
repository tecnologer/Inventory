package Inventory

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test"
)

var dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	DB_USER, DB_PASSWORD, DB_NAME)
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", dbinfo)
	checkErr(err)
}

func insertNewProduct(p Product) (int, error) {
	var lastInsertId int
	err := db.QueryRow("INSERT INTO Product(Id, Description, Price, Type, Quantity) VALUES($1,$2,$3, $4, $5) returning Id;",
		p.ID,
		p.Description,
		p.Price,
		p.Type,
		p.Qty,
	).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func getProducts() ([]Product, error) {
	var list []Product
	rows, err := db.Query("SELECT * FROM Product")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p = Product{}

		err = rows.Scan(&p.ID, &p.Description, &p.Price, &p.Type, &p.Qty)

		if err != nil {
			return nil, err
		}

		list = append(list, p)
	}

	return list, nil
}

func existsProduct(id int) (bool, error) {
	var exists bool
	rows, err := db.Query("SELECT EXISTS(SELECT 1 FROM product WHERE ID=$1)", id)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		err = rows.Scan(&exists)

		if err != nil {
			return false, err
		}
	}

	return exists, nil
}

func deleteProduct(id int) error {
	stmt, err := db.Prepare("DELETE FROM product WHERE ID=$1")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println(affect, "rows changed")
	return nil
}

func updateProduct(id int, desc string, price float32, productType Type) error {
	stmt, err := db.Prepare("UPDATE product SET Description=$2, Price=$3, Type=$4 WHERE ID=$1")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id, desc, price, productType)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println(affect, "rows changed")
	return nil
}

func addMovement(id int, qty float32) error {
	stmt, err := db.Prepare("UPDATE product SET Quantity=Quantity+$2 WHERE ID=$1")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id, qty)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println(affect, "rows changed")
	return nil
}

//CloseDB closes the current connection
func CloseDB() {
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
