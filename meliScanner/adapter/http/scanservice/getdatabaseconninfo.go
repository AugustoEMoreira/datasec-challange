package scanservice

import "github.com/AugustoEMoreira/datasec-challange/core/dto"

// GetDatabaseConnInfo implements domain.ScanService.
func (service service) GetDatabaseConnInfo(id string) (dbconnUrl *dto.DbConnectionInfo, err error) {
	db, err := service.GetDatabaseConnInfo(id)
	if err != nil {
		return nil, err
	}
	return db, nil
}
