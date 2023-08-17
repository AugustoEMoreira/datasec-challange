package scanusecase

import (
	"github.com/AugustoEMoreira/datasec-challange/core/domain"
	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

func (usecase usecase) Create(createScanRequest *dto.CreateScanRequest) (*domain.Datastorage, error) {
	ds, err := usecase.repository.Create(createScanRequest)
	if err != nil {
		return nil, err
	}
	return ds, nil
}
