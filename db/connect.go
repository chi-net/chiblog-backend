package db 

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"strconv"
	"github.com/chihuo2104/chiblog-backend/config"
)
func Connect *DB() {
	dsn := config.DatabaseUserName + ":" + config.DatabasePassword + "@tcp(" + config.DatabaseHost + ":" + strconv.Itoa(config.DatabasePort) + ")/" + config.DatabaseName
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