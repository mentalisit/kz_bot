package ReservCopy

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type ReservDB struct {
	db *sql.DB
}

func NewReservDB() *ReservDB {
	db, err := sql.Open("sqlite3", "hsbot/hsbot.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	if err = db.Ping(); err != nil {
		fmt.Println()
		return nil
	}
	return &ReservDB{db: db}
}
