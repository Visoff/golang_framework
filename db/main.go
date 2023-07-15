package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB DB_Class

func Connect() {
	db, err := sql.Open("postgres", "postgres://postgres:31415926@localhost:5432/dev?sslmode=disable")
	if err != nil {
		fmt.Println("Db error")
		fmt.Println(err)
		return
	}
	rows, err := db.Query(`select "id", "name" from "user"`)
	if err != nil {
		fmt.Println("Db error")
		fmt.Println(err)
		return
	}
	fmt.Println("connected db")
	defer rows.Close()
	DB = DB_Class{db}

	result := DB.Select("id", "name").From("user").Query()
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
