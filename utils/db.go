package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	dsn := "root:mysql@tcp(localhost:3306)/toyota_chatbot"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}
	return db, db.Ping()
}
