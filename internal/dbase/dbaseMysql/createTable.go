package dbaseMysql

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type CreateTable struct {
	db *sql.DB
}

func (c *CreateTable) AllTable() {

	c.table(users)
	c.table(sborkz)
	c.table(subscribe)
	c.table(config)
	c.table(numkz)
	c.table(rsevent)
	c.table(temptopevent)
	c.table(timer)
	//c.Config()
	//c.Numkz()
	//c.Rsevent()
	//c.Sborkz()
	//c.Subscribe()
	//c.Timer()
	//c.TempTop()
}
func (c *CreateTable) table(table string) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := c.db.ExecContext(ctx, table)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
}

/*
func (c *CreateTable) Config() {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := c.db.ExecContext(ctx, config)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
}
func (c *CreateTable) Numkz() {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := c.db.ExecContext(ctx, numkz)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
}
func (c *CreateTable) Rsevent() {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := c.db.ExecContext(ctx, rsevent)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
}
func (c *CreateTable) Sborkz() {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := c.db.ExecContext(ctx, sborkz)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
}
func (c *CreateTable) Subscribe() {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := c.db.ExecContext(ctx, subscribe)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
}

func (c *CreateTable) Timer() {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := c.db.ExecContext(ctx, timer)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
}

func (c *CreateTable) TempTop() {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := c.db.ExecContext(ctx, temptopevent)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
}

func (c *CreateTable) TEST() error {
	query := `CREATE TABLE IF NOT EXISTS product(
		product_id int primary key auto_increment,
		product_name text,
		product_price int,
		created_at datetime default CURRENT_TIMESTAMP,
		updated_at datetime default CURRENT_TIMESTAMP
		)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := c.db.ExecContext(ctx, query)
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


*/
