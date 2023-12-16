package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type note struct {
	Id         int64   `json:"id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Title      string  `json:"title"`
	Address    string  `json:"address"`
	Content    string  `json:"content"`
	UpdateTime int64   `json:"updateTime"`
	Visible    bool    `json:"visible"`

	UpdateTimeFormatted string `json:"updateTimeFormatted"`
}

var db *sql.DB

func initDB() {
	cfg := mysql.Config{
		User:   "imoonkin",
		Passwd: "passwd",
		Net:    "tcp",
		Addr:   "hostmysql:3306",
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
		"insert into sharedNote (latitude,longitude,title,address,content,updateTime,visible) values (?,?,?,?,?,?,?)",
		n.Latitude, n.Longitude, n.Title, n.Address, n.Content, time.Now().Unix(), true); err != nil {
		return err
	}
	return nil
}
func rangeFetch(lat1, lng1, lat2, lng2 float64) ([]note, error) {
	ret := []note{}
	rows, err := db.Query("select id,latitude,longitude,title,address,content,updateTime,visible from sharedNote where latitude between ? and ? and longitude between ? and ? limit 100", lat1, lat2, lng1, lng2)
	if err != nil {
		return ret, err
	}
	defer rows.Close()
	for rows.Next() {
		var n note
		if err = rows.Scan(&n.Id, &n.Latitude, &n.Longitude, &n.Title, &n.Address, &n.Content, &n.UpdateTime, &n.Visible); err != nil {
			return ret, err
		}
		ret = append(ret, n)
	}
	if rows.Err() != nil {
		return ret, rows.Err()
	}
	return ret, nil
}
