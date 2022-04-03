package db 

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)
func Connect *DB() {
	dsn := "root:171251@tcp(127.0.0.1:3306)/chiapp"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// log.Println("SQL连接成功")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}