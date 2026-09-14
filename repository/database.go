package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {
	dsn := "root:rootpassword@tcp(127.0.0.1:3306)/movie_store_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	ping := db.Ping()
	if ping != nil {
		log.Fatal("Error pinging the database:", ping)
	}
	return db

}
