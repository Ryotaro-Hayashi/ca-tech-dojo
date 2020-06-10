package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() (sqlHandler SqlHandler) {
	// db は、database/sqlパッケージの*DB型
	db, err := sql.Open("mysql", "root:rootpass@tcp(mysql:3306)/dojo_db")
	if err != nil {
		log.Panic("cannot connect db.")
	}

	sqlHandler = SqlHandler{}
	sqlHandler.Conn = db

	return
}