package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db    *sql.DB
	once  sync.Once
	local string = "host=localhost port=5432 user=saalcazar password=a1b2c3d4c0 dbname=surveys sslmode=disable"
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", local)
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}
		fmt.Println("Conectado a postgres")
	})
}

func Pool() *sql.DB {
	return db
}
