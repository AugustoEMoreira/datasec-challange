package scanrepository

import (
	"context"
	"fmt"

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
	query := `select host,port,password,user,db from datastorage where id = ?`
	rows, err := repository.db.QueryContext(ctx, query)
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
	dbc.UrlString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", queryResult.user, queryResult.password, queryResult.host, queryResult.port, queryResult.database)
	return &dbc, err
}
