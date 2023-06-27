package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:tonny14@tcp(127.0.0.1:3306)/goblog")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")

	sql := "INSERT INTO students(email, first_name, last_name) VALUES ('admin@gmail.com', 'admin','admin')"

	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)

	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
	fmt.Println("insterted")

	var (
		first_name string
		email      string
	)

	rows, err := db.Query("select first_name, email from students where id = ?", 4)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&first_name, &email)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(first_name, email)
		fmt.Println(email)
	}
}
