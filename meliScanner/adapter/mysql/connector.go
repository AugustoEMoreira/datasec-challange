package mysql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type PoolInterface interface {
	Close() error
	QueryContext(ctx context.Context, query string, arguments ...interface{}) (*sql.Rows, error)
	ExecContext(ctex context.Context, query string, arguments ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

func GetMainDBConnection(context context.Context) (*sql.Conn, error) {
	dbURL := viper.GetString("database.url")
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}

	conn, err := db.Conn(context)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
func GetCustomDBConnection(context context.Context, d string) (*sql.Conn, error) {

	fmt.Println(d)
	db, err := sql.Open("mysql", d)
	if err != nil {
		return nil, err
	}

	conn, err := db.Conn(context)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
