package scanrepository

import (
	"database/sql"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

func (repository repository) GetDbConnection(d *dto.DbConnectionInfo) (conn *sql.Conn, e error) {
	conn, err := repository.customDb.Conn("mysql", d.UrlString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
