package dbaseMysql

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func createTableConfig(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, config)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
	return nil
}

func createTableNumkz(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, numkz)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
	return nil
}

func createTableRsevent(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, rsevent)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}

	return nil
}

func createTableSborkz(db *sql.DB) error {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, sborkz)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}

	return nil
}

func createTableSubscribe(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, subscribe)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
	return nil
}

func createTableTimer(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, timer)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
	return nil
}

func createTableTempTop(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, temptopevent)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
	return nil
}

func createTableTEST(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS product(
		product_id int primary key auto_increment, 
		product_name text,
		product_price int,
		created_at datetime default CURRENT_TIMESTAMP, 
		updated_at datetime default CURRENT_TIMESTAMP
		)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
	return nil
}
