package mysql

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type CustomPoolInterface interface {
	Conn(driverName string, datasourceName string) (*sql.Conn, error)
	Close() error
	QueryContext(ctx context.Context, query string, arguments ...interface{}) (*sql.Rows, error)
	ExecContext(ctex context.Context, query string, arguments ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

func GetCustomConnection(context context.Context, connectionURL string) (*sql.Conn, error) {

	db, err := sql.Open("mysql", connectionURL)
	if err != nil {
		return nil, err
	}

	conn, err := db.Conn(context)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
