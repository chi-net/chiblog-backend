/*
 * @Author: chihuo2104
 * @Date: 2022-04-16 16:19:05
 * @LastEditTime: 2022-04-16 16:21:13
 * @LastEditors: chihuo2104
 * @Description: database connect with mysql
 * @FilePath: \chiblog-backend\db\connect.go
 * Powered by chihuo2104<chihuo2104@moekonnyaku.com>.All right reserved ©2018-now.
 */
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