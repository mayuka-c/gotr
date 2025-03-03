package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	"github.com/striversity/gotr/misc-ep006/ex03/streamer"
)

func main() {
	db, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
	if err != nil {
		logrus.Fatal(err)
	}
	query := `SELECT * FROM products;`

	dbStreamer := streamer.NewBuilder().WithQuery(query).WithInterval(1 * time.Second).WithDb(db).Build()

	dbStreamer.Start()
	time.Sleep(20 * time.Second)
	dbStreamer.Stop()
}
