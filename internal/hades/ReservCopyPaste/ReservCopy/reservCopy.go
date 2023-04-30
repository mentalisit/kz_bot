package ReservCopy

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

type ReservDB struct {
	db *sql.DB
}

func NewReservDB() *ReservDB {
	root, _ := os.Executable()
	curDir := filepath.Join(root, "..")
	dbPath := filepath.Join(curDir, "hsbot", "hsbot.db")

	db, err := sql.Open("sqlite3", dbPath)
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
