package scanusecase

import "github.com/AugustoEMoreira/datasec-challange/core/dto"

// SaveScan implements domain.ScanUseCase.
func (usecase usecase) SaveScan(sc *[]dto.ScanResult, dbId string) error {
	err := usecase.repository.SaveScan(sc, dbId)
	if err != nil {
		return err
	}
	return nil
}
