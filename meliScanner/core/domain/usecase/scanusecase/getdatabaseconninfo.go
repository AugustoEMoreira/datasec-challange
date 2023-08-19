package scanusecase

import "github.com/AugustoEMoreira/datasec-challange/core/dto"

// GetDatabaseConnInfo implements domain.ScanUseCase.
func (usecase usecase) GetDatabaseConnInfo(id string) (dbconnUrl *dto.DbConnectionInfo, err error) {
	dbc, err := usecase.repository.GetDatabaseConnInfo(id)
	if err != nil {
		return nil, err
	}
	return dbc, nil
}
