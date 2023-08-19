package scanusecase

import "github.com/AugustoEMoreira/datasec-challange/core/dto"

// SaveScan implements domain.ScanUseCase.
func (usecase usecase) SaveScan(sc *[]dto.ScanResult) error {
	err := usecase.repository.SaveScan(sc)
	if err != nil {
		return err
	}
	return nil
}
