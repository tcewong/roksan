package roksandb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func db() (db *sql.DB, err error) {
	return sql.Open("mysql", "root:P@ssw0rd@/Roksan")
}
