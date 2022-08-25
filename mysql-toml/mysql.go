package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

var db *sql.DB

func init_db() {

	var err error

	credential := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Host,
		strconv.Itoa(config.Mysql.Port),
		config.Mysql.Schema)

	db, err = sql.Open("mysql", credential)

	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(10)

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB Connect---------------------------------------------------------------------------", credential)
}
