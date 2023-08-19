package scanusecase

import (
	"database/sql"

	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

// GetDbConnection implements domain.ScanUseCase.
func (usecase usecase) GetDbConnection(dbc *dto.DbConnectionInfo) (conn *sql.Conn, err error) {
	connection, err := usecase.repository.GetDbConnection(dbc)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
