package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	//connect
	db, err := sql.Open(driverName, dataSourceName) //connection parameters for the database to connect to it
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	//test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &Adapter{db}, nil

}

func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db connection close: %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	//QUERY BUILDER
	queryString, args, err := sq.Insert("arith_history").Columns("date", "answers", "operation").Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	return nil
}
