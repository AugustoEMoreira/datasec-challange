package scanservice

import (
	"database/sql"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

// GetDbConnection implements domain.ScanService.
func (service service) GetDbConnection(db *dto.DbConnectionInfo) (conn *sql.Conn, err error) {
	connection, err := service.usecase.GetDbConnection(db)
	if err != nil {
		return nil, err
	}
	return connection, err
}
