package roksandb

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

// D contains details for inserting into db
type D struct {
	table string
	pairs map[string]string
}

type action func(db *sql.DB) (err error)

func db() (db *sql.DB, err error) {
	return sql.Open("mysql", "root:P@ssw0rd@/Roksan")
}

func process(act action) (err error) {
	db, err := sql.Open("mysql", "root:P@ssw0rd@/Roksan")
	if err == nil {
		defer db.Close()
		if act != nil {
			err = act(db)
		} else {
			err = errors.New("There is no any action")
		}
	}
	return
}

// Insert data into database
func Insert(d D) (err error) {
	var (
		keys   []string
		params []string
		vals   []interface{}
	)
	for key, val := range d.pairs {
		keys = append(keys, key)
		vals = append(vals, val)
		params = append(params, "?")
	}
	return process(
		func(db *sql.DB) (err error) {
			_, err = db.Exec("INSERT INTO "+d.table+"("+strings.Join(keys, ",")+") values("+strings.Join(params, ",")+");", vals...)
			return
		},
	)
}

// Query query the data from database
func Query(table, query string, shows ...string) (rows *sql.Rows, err error) {
	s := "*"
	if len(shows) > 0 {
		s = strings.Join(shows, ",")
	}
	command := "SELECT " + s + " FROM " + table
	if query != "" {
		command += " WHERE " + query
	}
	err = process(
		func(db *sql.DB) (err error) {
			rows, err = db.Query(command)
			return
		},
	)
	return
}

func Test() {
	err := DeleteShowrooms("Testing1")
	if err != nil {
		panic(err)
	}
}
