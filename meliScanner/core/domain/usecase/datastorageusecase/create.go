package datastorageusecase

import (
	"github.com/AugustoEMoreira/datasec-challange/core/domain"
	"github.com/AugustoEMoreira/datasec-challange/core/dto"
)

func (usecase usecase) Create(datastorageRequest *dto.CreateDatastoreRequest) (*domain.Datastorage, error) {
	ds, err := usecase.repository.Create(datastorageRequest)
	if err != nil {
		return nil, err
	}
	return ds, nil
}
