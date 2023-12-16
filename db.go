package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type note struct {
	id         int64
	latitude   float64
	longitude  float64
	title      string
	address    string
	content    string
	updateTime int64
	visible    bool

	updateTimeFormatted string
}

var db *sql.DB

func initDB() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "passwd",
		Net:    "tcp",
		Addr:   "host.k3d.internal:3306",
		DBName: "sharednotedb",
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
		"insert into sharednote (latitude,longitude,title,address,content,updateTime,visible) values (?,?,?,?,?,?,?)",
		n.latitude, n.longitude, n.title, n.address, n.content, time.Now().Unix(), true); err != nil {
		return err
	}
	return nil
}
func rangeFetch(lat1, lng1, lat2, lng2 float64) ([]note, error) {
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
