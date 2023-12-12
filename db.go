package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type note struct {
	latitude  float32
	longitude float32
	title     string
	address   string
	content   string
}

var db *sql.DB

func initDB() {
	cfg := mysql.Config{
		User:   "owen",
		Passwd: "123haha123",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "sharedNotes",
	}
	var err error
	if db, err = sql.Open("mysql", cfg.FormatDSN()); err != nil {
		log.Panicln(err)
	}
	if err = db.Ping(); err != nil {
		log.Panicln(err)
	}
	log.Println("db connected")

}

func addNote(n note) error {
	if _, err := db.Exec(
		"insert into sharednote (latitude,longitude,title,address,content) values (?,?,?,?,?)",
		n.latitude, n.longitude, n.title, n.address, n.content); err != nil {
		return err
	}
	return nil
}
func rangeFetch(lat1, lng1, lat2, lng2 float32) ([]note, error) {
	ret := []note{}
	rows, err := db.Query("select * from sharednote where latitude between lat1 and lat2 and longitude between lng1 and lng2 limit 100")
	if err != nil {
		return ret, err
	}
	defer rows.Close()
	for rows.Next() {
		var n note
		if err = rows.Scan(&n.latitude, &n.longitude, &n.title, &n.address, &n.content); err != nil {
			return ret, err
		}
		ret = append(ret, n)
	}
	if rows.Err() != nil {
		return ret, rows.Err()
	}
	return ret, nil
}
