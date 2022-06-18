package dbaseMysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbname   = "kzbot"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func (d *Db) DbConnection() *sql.DB {
	var err error
	d.Db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		//return nil, err
	}
	d.Db.SetMaxOpenConns(10)
	d.Db.SetMaxIdleConns(10)
	d.Db.SetConnMaxLifetime(time.Minute * 5)
	return d.Db
	/*
		//ctx, cancelfunc = context.WithTimeout(context.Background(), 7*time.Second)
		//defer cancelfunc()
		//err = db.PingContext(ctx)
		//if err != nil {
		//	log.Printf("Errors %s pinging DB", err)
			//return nil, err
		//}
		//log.Printf("Connected to DB %s successfully\n", dbname)
		//if no == 1 {
			err = createTableConfig(db)
			err = createTableNumkz(db)
			err = createTableRsevent(db)
			err = createTableSborkz(db)
			err = createTableSubscribe(db)
			err = createTableTimer(db)
			err = createTableTempTop(db)
			fmt.Println("Таблицы созданы ошибок вроде нет ")
			fmt.Println(err)
		}
	*/
	//NewDb(Db)
	//d.Db = Db
}

/*
func init() {
	db, err := DbConnection()
	if err != nil {
		log.Printf("Error %s when getting dbaseMysql connection", err)
		return
	}
	defer db.Close()
	if err != nil {
		log.Printf("Create product table failed with error %s", err)
		return
	}
}

*/
