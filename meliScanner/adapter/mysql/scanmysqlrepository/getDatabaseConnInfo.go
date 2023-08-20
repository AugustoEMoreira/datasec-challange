package scanrepository

import (
	"context"
	"fmt"
	"net/url"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

type queryResult struct {
	host     string
	port     string
	password string
	user     string
	database string
}

// GetDatabaseConnInfo implements domain.ScanRepository.
func (repository repository) GetDatabaseConnInfo(id string) (dbconnUrl *dto.DbConnectionInfo, err error) {
	ctx := context.Background()
	query := `select host,port,pass,user,db from datastorage where id = ?`
	stmt, err := repository.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var queryResult queryResult
	var dbc dto.DbConnectionInfo
	for rows.Next() {
		err := rows.Scan(&queryResult.host, &queryResult.port, &queryResult.password, &queryResult.user, &queryResult.database)
		if err != nil {
			return nil, err
		}
	}

	escapepassword := url.QueryEscape(queryResult.password)
	dbc.UrlString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", queryResult.user, escapepassword, queryResult.host, queryResult.port, queryResult.database)
	return &dbc, err
}
