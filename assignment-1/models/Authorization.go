package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Name, pass, dbname string

func GetDB() (*sql.DB, error) {
	//fmt.Println("Get access to database")
	//fmt.Print("name: ")
	//fmt.Scan(&Name)
	//fmt.Print("password: ")
	//fmt.Scan(&pass)
	//fmt.Print("Database name: ")
	//fmt.Scan(&dbname)
	//fmt.Println()
	//conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", Name, pass, dbname)
	conn := fmt.Sprintf("user=postgres password=Merei04977773@ dbname=shop sslmode=disable")

	db, err := sql.Open("postgres", conn)
	return db, err
}
